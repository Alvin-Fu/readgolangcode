// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import (
	"readruntime/internal/sys"
	"unsafe"
)

type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

// 不在GC堆上分配内存，常用在运行时用作低层次的内部结构，避免调度器和内存分配中的写屏障
// An notInHeapSlice is a slice backed by go:notinheap memory.
type notInHeapSlice struct {
	array *notInHeap
	len   int
	cap   int
}

// maxElems is a lookup table containing the maximum capacity for a slice.
// The index is the size of the slice element.
// 这个是元素大小和切片大小的对照表
var maxElems = [...]uintptr{
	^uintptr(0),
	maxAlloc / 1, maxAlloc / 2, maxAlloc / 3, maxAlloc / 4,
	maxAlloc / 5, maxAlloc / 6, maxAlloc / 7, maxAlloc / 8,
	maxAlloc / 9, maxAlloc / 10, maxAlloc / 11, maxAlloc / 12,
	maxAlloc / 13, maxAlloc / 14, maxAlloc / 15, maxAlloc / 16,
	maxAlloc / 17, maxAlloc / 18, maxAlloc / 19, maxAlloc / 20,
	maxAlloc / 21, maxAlloc / 22, maxAlloc / 23, maxAlloc / 24,
	maxAlloc / 25, maxAlloc / 26, maxAlloc / 27, maxAlloc / 28,
	maxAlloc / 29, maxAlloc / 30, maxAlloc / 31, maxAlloc / 32,
}

// maxSliceCap returns the maximum capacity for a slice.
// 返回一个切片的最大容量， 根据元素的大小
func maxSliceCap(elemsize uintptr) uintptr {
	if elemsize < uintptr(len(maxElems)) {
		return maxElems[elemsize]
	}
	return maxAlloc / elemsize
}

// panic 切片的长度越界
func panicmakeslicelen() {
	panic(errorString("makeslice: len out of range"))
}

// panic 切片的容量越界
func panicmakeslicecap() {
	panic(errorString("makeslice: cap out of range"))
}

// 创建一个切片， et表示一个元素的类型，即需要创建一个什么样的切片， 后面的是长度和容量
func makeslice(et *_type, len, cap int) slice {
	// NOTE: The len > maxElements check here is not strictly necessary,
	// but it produces a 'len out of range' error instead of a 'cap out of range' error when someone does make([]T, bignumber).
	// 'cap out of range' is true too, but since the cap is only being supplied implicitly, saying len is clearer.
	// See issue 4085.
	// 返回这类型元素可分配的切片的最大容量
	maxElements := maxSliceCap(et.size)
	// 这个检查其实不是必需的
	// 先需要检查的是长度，因为切片的长度需要比容量小
	if len < 0 || uintptr(len) > maxElements {
		panicmakeslicelen()
	}

	if cap < len || uintptr(cap) > maxElements {
		panicmakeslicecap()
	}
	// 分配空间
	p := mallocgc(et.size*uintptr(cap), et, true)
	return slice{p, len, cap}
}

// 使用了类型转换为什么还要这个方法
func makeslice64(et *_type, len64, cap64 int64) slice {
	len := int(len64)
	if int64(len) != len64 {
		panicmakeslicelen()
	}

	cap := int(cap64)
	if int64(cap) != cap64 {
		panicmakeslicecap()
	}

	return makeslice(et, len, cap)
}

// growslice handles slice growth during append.动态增长切片
// It is passed the slice element type, the old slice, and the desired new minimum capacity,
// and it returns a new slice with at least that capacity, with the old data copied into it.
// 这段讲的是这个方法是根据元素的类型，老的切片，新的需要的最少的容量， 然后分配一个新的，将老的copy到新的中
// The new slice's length is set to the old slice's length,
// NOT to the new requested capacity.
// This is for codegen convenience. The old slice's length is used immediately
// to calculate where to write new values during an append.
// 在写入得时候这个容量会被重新估计， 为了性能
// 在运行过程中切片不够时自动增长
//
// TODO: When the old backend is gone, reconsider this decision.
// The SSA backend might prefer the new length or to return only ptr/cap and save stack space.
// 三个参数，et是元素类型，old是老的切片， cap是需要分配的容量，即指定的
func growslice(et *_type, old slice, cap int) slice {
	if raceenabled {
		callerpc := getcallerpc()
		racereadrangepc(old.array, uintptr(old.len*int(et.size)), callerpc, funcPC(growslice))
	}
	if msanenabled {
		msanread(old.array, uintptr(old.len*int(et.size)))
	}

	if et.size == 0 {
		if cap < old.cap {
			panic(errorString("growslice: cap out of range"))
		}
		// append should not create a slice with nil pointer but non-zero len.
		// We assume that append doesn't need to preserve old.array in this case.
		return slice{unsafe.Pointer(&zerobase), old.len, cap}
	}

	//下面的代码就是，设置需要扩容的大小
	newcap := old.cap
	// 两倍之前的容量
	doublecap := newcap + newcap
	if cap > doublecap {
		newcap = cap
	} else {
		if old.len < 1024 {
			newcap = doublecap
		} else {
			// Check 0 < newcap to detect overflow and prevent an infinite loop.
			// 检查溢出，并且使得newcap尽可能的大一点
			for 0 < newcap && newcap < cap {
				newcap += newcap / 4
			}
			// Set newcap to the requested cap when the newcap calculation overflowed.
			// 当老的切片溢出是，分配的容量为，指定的cap
			if newcap <= 0 {
				newcap = cap
			}
		}
	}

	var overflow bool
	// 这三个参数分别是老的切片的大小， 新的切片大小，容量
	var lenmem, newlenmem, capmem uintptr
	// Specialize for common values of et.size.
	// For 1 we don't need any division/multiplication.
	// For sys.PtrSize, compiler will optimize division/multiplication into a shift by a constant.
	// For powers of 2, use a variable shift.
	// 对于不同大小的元素的扩容方式不同，主要是将乘法变成移位操作
	switch {
	case et.size == 1:
		lenmem = uintptr(old.len)
		newlenmem = uintptr(cap)
		capmem = roundupsize(uintptr(newcap))
		overflow = uintptr(newcap) > maxAlloc
		newcap = int(capmem)
	case et.size == sys.PtrSize:
		lenmem = uintptr(old.len) * sys.PtrSize
		newlenmem = uintptr(cap) * sys.PtrSize
		capmem = roundupsize(uintptr(newcap) * sys.PtrSize)
		overflow = uintptr(newcap) > maxAlloc/sys.PtrSize
		newcap = int(capmem / sys.PtrSize)
	case isPowerOfTwo(et.size):
		var shift uintptr
		if sys.PtrSize == 8 {
			// Mask shift for better code generation.
			shift = uintptr(sys.Ctz64(uint64(et.size))) & 63
		} else {
			shift = uintptr(sys.Ctz32(uint32(et.size))) & 31
		}
		lenmem = uintptr(old.len) << shift
		newlenmem = uintptr(cap) << shift
		capmem = roundupsize(uintptr(newcap) << shift)
		overflow = uintptr(newcap) > (maxAlloc >> shift)
		newcap = int(capmem >> shift)
	default:
		lenmem = uintptr(old.len) * et.size
		newlenmem = uintptr(cap) * et.size
		capmem = roundupsize(uintptr(newcap) * et.size)
		overflow = uintptr(newcap) > maxSliceCap(et.size)
		newcap = int(capmem / et.size)
	}

	// The check of overflow (uintptr(newcap) > maxSliceCap(et.size))
	// in addition to capmem > _MaxMem is needed to prevent(防止) an overflow(溢出)
	// which can be used to trigger(触发) a segfault(段错误) on 32bit architectures(架构)
	// with this example program:
	//
	// type T [1<<27 + 1]int64
	//
	// var d T
	// var s []T
	//
	// func main() {
	//   s = append(s, d, d, d, d)
	//   print(len(s), "\n")
	// }
	if cap < old.cap || overflow || capmem > maxAlloc {
		panic(errorString("growslice: cap out of range"))
	}

	var p unsafe.Pointer
	if et.kind&kindNoPointers != 0 {
		// 进行内存分配
		p = mallocgc(capmem, nil, false)
		// 将老的数组数组拷贝给新的数组，个数为lenmem
		memmove(p, old.array, lenmem)
		// The append() that calls growslice is going to overwrite from old.len to cap (which will be the new length).
		// Only clear the part that will not be overwritten.
		// 清理没有数据的那部分就是长度到容量的那部分
		memclrNoHeapPointers(add(p, newlenmem), capmem-newlenmem)
	} else {
		// Note: can't use rawmem (which avoids zeroing of memory避免内存归零),
		// because then GC can scan uninitialized(未初始化) memory.
		p = mallocgc(capmem, et, true)
		if !writeBarrier.enabled {
			memmove(p, old.array, lenmem)
		} else {
			for i := uintptr(0); i < lenmem; i += et.size {
				typedmemmove(et, add(p, i), add(old.array, i))
			}
		}
	}

	return slice{p, old.len, newcap}
}

func isPowerOfTwo(x uintptr) bool {
	return x&(x-1) == 0
}

// 切片copy
func slicecopy(to, fm slice, width uintptr) int {
	if fm.len == 0 || to.len == 0 {
		return 0
	}

	n := fm.len
	if to.len < n {
		n = to.len
	}

	if width == 0 {
		return n
	}

	if raceenabled {
		callerpc := getcallerpc()
		pc := funcPC(slicecopy)
		racewriterangepc(to.array, uintptr(n*int(width)), callerpc, pc)
		racereadrangepc(fm.array, uintptr(n*int(width)), callerpc, pc)
	}
	if msanenabled {
		msanwrite(to.array, uintptr(n*int(width)))
		msanread(fm.array, uintptr(n*int(width)))
	}

	size := uintptr(n) * width
	if size == 1 { // common case worth about 2x to do here
		// TODO: is this still worth it with new memmove impl?
		*(*byte)(to.array) = *(*byte)(fm.array) // known to be a byte pointer
	} else {
		memmove(to.array, fm.array, size)
	}
	return n
}

func slicestringcopy(to []byte, fm string) int {
	if len(fm) == 0 || len(to) == 0 {
		return 0
	}

	n := len(fm)
	if len(to) < n {
		n = len(to)
	}

	if raceenabled {
		callerpc := getcallerpc()
		pc := funcPC(slicestringcopy)
		racewriterangepc(unsafe.Pointer(&to[0]), uintptr(n), callerpc, pc)
	}
	if msanenabled {
		msanwrite(unsafe.Pointer(&to[0]), uintptr(n))
	}

	memmove(unsafe.Pointer(&to[0]), stringStructOf(&fm).str, uintptr(n))
	return n
}
