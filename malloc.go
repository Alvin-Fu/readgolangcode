// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Memory allocator.
// 内存分配器
// This was originally based on tcmalloc, but has diverged quite a bit.
// http://goog-perftools.sourceforge.net/doc/tcmalloc.html
//
// 内存管理单元和objects的关系，内存管理器会将内存分成大约70个类，每个类都有一组相同大小的对象，
// 每个页都可以分成一个大小类的对象集使用空闲位图（bitmap）管理器
// The main allocator works in runs of pages.
// Small allocation sizes (up to and including 32 kB) are
// rounded to one of about 70 size classes, each of which
// has its own free set of objects of exactly that size.
// Any free page of memory can be split into a set of objects
// of one size class, which are then managed using a free bitmap.
//
// 分配器的数据结构
// The allocator's data structures are:
//
//	fixalloc:
// 一个固定大小的heap对象的空闲分配器，用来管理分配器的存储
// a free-list allocator for fixed-size off-heap objects, used to manage storage used by the allocator.
//  堆分配器，管理的粒度是8192byte
//	mheap: the malloc heap, managed at page (8192-byte) granularity.
//  有mheap管理的一系列页面
//	mspan: a run of pages managed by the mheap.
//	所有给定大小的spans的集合，可以将mcentral理解成一个spans的管理池
//	mcentral: collects all spans of a given size class.
//	每个p上的缓存
//	mcache: a per-P cache of mspans with free space.
//  分配器的统计信息
//	mstats: allocation statistics.

//	在go中有三个内存分配层级即大中小
// 分配一个小对象的时候会进入缓存结构
// Allocating a small object proceeds up a hierarchy of caches:
//
//  下面的意思是如果可以在当前的空闲位图中找到span就直接分配并且该过程不需要加锁
//	1. Round the size up to one of the small size classes
//	   and look in(查找) the corresponding(相应) mspan in this P's mcache.
//	   Scan(扫描) the mspan's free bitmap to find a free slot(位置).
//	   If there is a free slot, allocate(分配) it.
//	   This can all be done without acquiring a lock.这些操作多可以在不获得锁
//
//  如果第一步如果没有则向mcentral中申请一个mspans列表
//	2. If the mspan has no free slots, obtain a new mspan
//	   from the mcentral's list of mspans of the required size class that have free space.
//	   Obtaining a whole span amortizes the cost of locking the mcentral. 这个操作是需要加锁的
//	如果第二步还是没有获取到就从mheap中获取
//	3. If the mcentral's mspan list is empty, obtain a run
//	   of pages from the mheap to use for the mspan.
//
//	如果mheap中没有获取到就向操作系统申请，不少于1MB. 一次性拿出一个大的内存可以减少和操作系统交互的成本
//	4. If the mheap is empty or has no page runs large enough,
//	   allocate a new group of pages (at least 1MB) from the operating system.
// 	   Allocating a large run of pages amortizes the cost of talking to the operating system.
//
// 扫描mspan并释放objects是一个类似的层级
// Sweeping an mspan and freeing objects on it proceeds up a similar hierarchy:
//	下面介绍的就是：如果在释放的过程中本层的的空闲过多的时候会将一部分归还给上层保证整体内存占用的平衡
//	1. If the mspan is being swept in response to allocation, it
//	   is returned to the mcache to satisfy the allocation.
//
//	2. Otherwise, if the mspan still has allocated objects in it,
//	   it is placed on the mcentral free list for the mspan's size
//	   class.
//
//	3. Otherwise, if all objects in the mspan are free, the mspan
//	   is now "idle", so it is returned to the mheap and no longer
//	   has a size class.
//	   This may coalesce it with adjacent idle mspans.
//
//	4. If an mspan remains idle for long enough, return its pages
//	   to the operating system.
//
// 分配一个大的对象的时候会直接在mheap上分配
// Allocating and freeing a large object uses the mheap
// directly, bypassing(绕过) the mcache and mcentral.
//
// 仅当mspan.needzero为false时，mspan中的自由对象槽才会归零。如果needzero为true，则对象在分配时归零。以这种方式延迟归零有很多好处
// Free object slots in an mspan are zeroed only if mspan.needzero is false.
// If needzero is true, objects are zeroed as they are allocated.
// There are various benefits to delaying zeroing this way:
//
//	1. Stack frame allocation can avoid zeroing altogether.
//
//	2. It exhibits better temporal locality, since the program is
//	   probably about to write to the memory.
//
//	3. We don't zero pages that never get reused.
// 1.堆栈帧分配可以完全避免归零。
// 2.它表现出更好的时间局部性，因为程序可能要写入内存。
// 3.我们不会零页面永远不会被重复使用。
//
//
// Virtual memory layout 虚拟内存的设计
//
// 堆是由一组arena组成
// The heap consists of a set of arenas, which are 64MB on 64-bit and 4MB on 32-bit (heapArenaBytes).
// Each arena's start address is also aligned to the arena size. 每一组arena开始的地址也是arena的大小
//
// Each arena has an associated heapArena object that stores the metadata for that arena:
// the heap bitmap for all words in the arena and the span map for all pages in the arena.
// heapArena objects are themselves allocated off-heap.
//
// Since arenas are aligned, the address space can be viewed as a series of arena frames.
// The arena map (mheap_.arenas) maps from arena frame number to *heapArena, or nil for parts of the address space not backed by the Go heap.
// The arena map is structured as a two-level array consisting of a "L1" arena map and many "L2" arena maps;
// however, since arenas are large, on many architectures, the arena map consists of a single, large L2 map.
//
// The arena map covers the entire possible address space, allowing the Go heap to use any part of the address space.
// The allocator attempts to keep arenas contiguous so that large spans (and hence large objects) can cross arenas.

package runtime

import (
	"readruntime/internal/atomic"
	"readruntime/internal/sys"
	"unsafe"
)

const (
	debugMalloc = false

	maxTinySize   = _TinySize
	tinySizeClass = _TinySizeClass
	maxSmallSize  = _MaxSmallSize

	pageShift = _PageShift
	pageSize  = _PageSize
	pageMask  = _PageMask
	// By construction, single page spans of the smallest object class have the most objects per span.
	// 通过构造，最小对象类的单页spans中有最多的对象
	maxObjsPerSpan = pageSize / 8

	mSpanInUse = _MSpanInUse

	concurrentSweep = _ConcurrentSweep
	// go中内存页的大小， 8k
	_PageSize = 1 << _PageShift
	_PageMask = _PageSize - 1

	// _64bit = 1 on 64-bit systems, 0 on 32-bit systems
	_64bit = 1 << (^uintptr(0) >> 63) / 2

	// Tiny allocator parameters, see "Tiny allocator" comment in malloc.go.小分配器的参数
	_TinySize      = 16
	_TinySizeClass = int8(2)

	_FixAllocChunk = 16 << 10               // Chunk size for FixAlloc fixalloc块的大小
	_MaxMHeapList  = 1 << (20 - _PageShift) // Maximum page length for fixed-size list in MHeap.MHeap中固定大小列表的最大页长

	// Per-P, per order stack segment cache size.每个p，每个订单栈段缓存的大小
	_StackCacheSize = 32 * 1024

	// Number of orders that get caching. 获得缓存的订单数量
	// Order 0 is FixedStack and each successive order is twice as large.
	// We want to cache 2KB, 4KB, 8KB, and 16KB stacks.
	// Larger stacks will be allocated directly.较大的堆栈将直接分配
	// Since FixedStack is different on different systems, we must vary NumStackOrders to keep the same maximum cached size.
	// 虽然不同系统之间的fixedStack是不同的，但是我们需要保持相同的cached的大小
	//   OS               | FixedStack | NumStackOrders
	//   -----------------+------------+---------------
	//   linux/darwin/bsd | 2KB        | 4
	//   windows/32       | 4KB        | 3
	//   windows/64       | 8KB        | 2
	//   plan9            | 4KB        | 3
	_NumStackOrders = 4 - sys.PtrSize/4*sys.GoosWindows - 1*sys.GoosPlan9

	// heapAddrBits is the number of bits in a heap address.
	// heapAddrBits是堆地址的bits数
	// On amd64, addresses are sign-extended beyond heapAddrBits.
	// On other arches, they are zero-extended.
	//
	// On 64-bit platforms(平台), we limit this to 48 bits based on(基于) a combination(结合) of hardware(硬件) and OS limitations(限制).
	// 在64位平台中，我们基于硬件和系统的因素我们把这个限制在48位
	// amd64 hardware limits addresses to 48 bits, sign-extended(符号扩展) to 64 bits.
	// 硬件的限制是48位但是符号的扩展是64位
	// Addresses where the top 16 bits are not either all 0 or all 1 are "non-canonical"(非标准的) and invalid(无效的).
	// 地址的前16位是无效的，即不被使用的
	// Because of these "negative(否定的)" addresses,
	// we offset addresses by 1<<47 (arenaBaseOffset) on amd64 before computing indexes into the heap arenas index.
	// In 2017, amd64 hardware added support for 57 bit addresses;
	// however, currently(目前) only Linux supports this extension(延长) and the kernel(内核) will never choose an
	// address above(超过) 1<<47 unless mmap is called with a hint address above 1<<47 (which we never do).
	//这里就是说不能超过48位，虽然linux已经延长到了57位，但是还是用不了
	// arm64 hardware (as of ARMv8) limits user addresses to 48 bits, in the range [0, 1<<48).
	// 在arm64硬件中限制使用地址是48位，即[0, 1<<48)
	// ppc64, mips64, and s390x support arbitrary(任意的) 64 bit addresses in hardware.
	// 这些平台支持的是全部的64位
	// However, since Go only supports Linux on these, we lean on OS limits.
	// 但是由于go只支持linux，因此我们依赖于操作系统的限制
	// Based on Linux's processor.h(可以看看内核的这个文件), the user address space is limited as follows(遵循) on 64-bit architectures(架构):
	//
	// Architecture  Name              Maximum Value (exclusive)
	// ---------------------------------------------------------------------
	// amd64         TASK_SIZE_MAX     0x007ffffffff000 (47 bit addresses)
	// arm64         TASK_SIZE_64      0x01000000000000 (48 bit addresses)
	// ppc64{,le}    TASK_SIZE_USER64  0x00400000000000 (46 bit addresses)
	// mips64{,le}   TASK_SIZE64       0x00010000000000 (40 bit addresses)
	// s390x         TASK_SIZE         1<<64 (64 bit addresses)
	//
	// These limits may increase over time, but are currently at most 48 bits except on s390x.
	// 随着时间的推移这些限制可能会增加，但是目前是48位，s390x除外
	// On all architectures, Linux starts placing mmap'd regions at addresses that are significantly below 48 bits,
	// 在所有的架构中，Linux开始时将放在mmap’d的地址放在低于48位地址上
	// so even if it's possible to exceed Go's 48 bit limit, it's extremely unlikely in practice.
	// 因此即使超过了go的48位的限制，这也是在实践中不可能的
	// On 32-bit platforms, we accept the full 32-bit address space because doing so is cheap.
	// 在32为体系中，接受了完整的地址空间，这样是简单的
	// mips32 only has access to the low 2GB of virtual memory, so we further limit it to 31 bits.
	// mips32
	// WebAssembly currently has a limit of 4GB linear memory.
	heapAddrBits = (_64bit*(1-sys.GoarchWasm))*48 + (1-_64bit+sys.GoarchWasm)*(32-(sys.GoarchMips+sys.GoarchMipsle))

	// maxAlloc is the maximum size of an allocation(分配).
	// On 64-bit, it's theoretically(理论上) possible to allocate 1<<heapAddrBits bytes. On
	// 32-bit, however, this is one less than 1<<32 because the
	// number of bytes in the address space doesn't actually fit
	// in a uintptr.
	maxAlloc = (1 << heapAddrBits) - (1-_64bit)*1

	// The number of bits in a heap address, the size of heap
	// arenas, and the L1 and L2 arena map sizes are related(关联) by
	//
	//   (1 << addrBits) = arenaBytes * L1entries * L2entries
	//
	// Currently, we balance these as follows:
	//
	//       Platform  Addr bits  Arena size  L1 entries  L2 size
	// --------------  ---------  ----------  ----------  -------
	//       */64-bit         48        64MB           1     32MB
	// windows/64-bit         48         4MB          64      8MB
	//       */32-bit         32         4MB           1      4KB
	//     */mips(le)         31         4MB           1      2KB

	// heapArenaBytes is the size of a heap arena. The heap
	// consists of mappings of size heapArenaBytes, aligned to
	// heapArenaBytes. The initial heap mapping is one arena.
	//
	// This is currently 64MB on 64-bit non-Windows and 4MB on
	// 32-bit and on Windows. We use smaller arenas on Windows
	// because all committed memory is charged to the process,
	// even if it's not touched. Hence, for processes with small
	// heaps, the mapped arena space needs to be commensurate.
	// This is particularly(格外，异常) important with the race detector(种类探测),
	// since it significantly amplifies(放大，扩大) the cost of committed memory.
	heapArenaBytes = 1 << logHeapArenaBytes

	// logHeapArenaBytes is log_2 of heapArenaBytes. For clarity(为清楚起见),
	// prefer using heapArenaBytes where possible (we need the
	// constant to compute(计算) some other constants).
	logHeapArenaBytes = (6+20)*(_64bit*(1-sys.GoosWindows)) + (2+20)*(_64bit*sys.GoosWindows) + (2+20)*(1-_64bit)

	// heapArenaBitmapBytes is the size of each heap arena's bitmap.bitmap的大小
	heapArenaBitmapBytes = heapArenaBytes / (sys.PtrSize * 8 / 2)

	pagesPerArena = heapArenaBytes / pageSize

	// arenaL1Bits is the number of bits of the arena number covered by the first level arena map.
	//
	// This number should be small, since the first level arena map requires PtrSize*(1<<arenaL1Bits) of space in the binary's BSS.
	// It can be zero, in which case the first level index is effectively unused.
	// There is a performance benefit to this, since the generated code can be more efficient, but comes at the cost of having a large L2 mapping.
	//
	// We use the L1 map on 64-bit Windows because the arena size is small, but the address space is still 48 bits, and there's a high cost to having a large L2.
	arenaL1Bits = 6 * (_64bit * sys.GoosWindows) // L1的大小

	// arenaL2Bits is the number of bits of the arena number covered by the second level arena index.
	//
	// The size of each arena map allocation is proportional to 1<<arenaL2Bits, so it's important that this not be too large.
	// 48 bits leads to 32MB arena index allocations(内存分配), which is about the practical threshold(实际的阈值).
	arenaL2Bits = heapAddrBits - logHeapArenaBytes - arenaL1Bits

	// arenaL1Shift is the number of bits to shift an arena frame
	// number by to compute an index into the first level arena map.
	arenaL1Shift = arenaL2Bits

	// arenaBits is the total bits in a combined arena map index.这是所有的arena映射索引的所有位
	// This is split between the index into the L1 arena map and the L2 arena map.
	arenaBits = arenaL1Bits + arenaL2Bits

	// arenaBaseOffset is the pointer value that corresponds to(相对应的)
	// index 0 in the heap arena map.
	//
	// On amd64, the address space is 48 bits, sign extended(符号扩展) to 64
	// bits. This offset lets us handle "negative" addresses (or
	// high addresses if viewed as unsigned).
	//
	// On other platforms(平台), the user address space is contiguous(相连的)
	// and starts(开始) at 0, so no offset(偏移量) is necessary.
	arenaBaseOffset uintptr = sys.GoarchAmd64 * (1 << 47)

	// Max number of threads to run garbage collection.表示最大可以跑gc的线程数
	// 2, 3, and 4 are all plausible(看似合理的) maximums depending on the hardware details of the machine.
	// 这些看是合理的数量是依赖与计算机的硬件基础的
	// The garbage collector scales well to 32 cpus.
	// 这个gc可以平衡到最好的是32cpu
	_MaxGcproc = 32

	// minLegalPointer is the smallest possible legal pointer.
	// This is the smallest possible architectural page size, since we assume(假设) that the first page is never mapped(映射).
	// 这可能是最小架构页的大小，我们假设从来没有映射过第一个页。
	// This should agree with minZeroPage in the compiler.
	// 这应该和编译器的minZeroPage是一致的
	minLegalPointer uintptr = 4096
)

// physPageSize is the size in bytes of the OS's physical pages.  这是系统的物理页的大小
// Mapping and unmapping operations must be done at multiples(倍数) of physPageSize.
//
// This must be set by the OS init code (typically in osinit) before mallocinit.
var physPageSize uintptr

// OS-defined helpers: 系统定义的一些帮助函数
//
// sysAlloc obtains(获得) a large chunk of zeroed memory from the
// operating system, typically on the order of a hundred kilobytes
// or a megabyte.
// NOTE: sysAlloc returns OS-aligned memory, but the heap allocator
// may use larger alignment, so the caller must be careful to realign(重新对齐) the
// memory obtained by sysAlloc.
//
// SysUnused notifies(通知) the operating system that the contents of the memory region are no longer needed and can be reused for other purposes(用途).
// SysUnused是用于通知操作系统这块内存不在使用了，可以回收了
// SysUsed notifies the operating system that the contents of the memory region are needed again.
// SysUsed是用于通知操作系统这块内存将会被再次使用
//
// SysFree returns it unconditionally(无条件的),
// this is only used if an out-of-memory error has been detected midway through an allocation.
// 仅仅是在分配的过程中发现内存不足时的使用
// It is okay if SysFree is a no-op.
//
// SysReserve reserves address space without allocating memory.
// 用于保留地址空间，而不分配内存
// If the pointer passed to it is non-nil, the caller wants the reservation there,
// but SysReserve can still choose another location if that one is unavailable.
// NOTE: SysReserve returns OS-aligned memory,
// but the heap allocator may use larger alignment, so the caller must be careful to realign the memory obtained by sysAlloc.
//
// SysMap maps previously reserved address space for use.
// 映射一段预留的的地址空间， os不一定会给
//
// SysFault marks a (already sysAlloc'd) region to fault if accessed. Used only for debugging the runtime.
//

// 实现内存分配的一些初始化，检查对象规格大小对照表，还将连续的虚拟地址划分为三块
//
/*
512MB      16GB            512GB
+-------+-------------+-------------------+
| spans |    bitmap   |       arena       |
+-------+-------------+-------------------+
*/
// 这三个区域是按需进行扩展，不需要预先分配
func mallocinit() {
	if class_to_size[_TinySizeClass] != _TinySize {
		throw("bad TinySizeClass")
	}

	testdefersizes()
	// 检查大小要是2的次方
	if heapArenaBitmapBytes&(heapArenaBitmapBytes-1) != 0 {
		// heapBits expects modular arithmetic on bitmap addresses to work.
		throw("heapArenaBitmapBytes not a power of 2")
	}

	// Copy class sizes out for statistics table.复制统计表中的类大小
	for i := range class_to_size {
		memstats.by_size[i].size = uint32(class_to_size[i])
	}

	// Check physPageSize. 检查物理页的大小
	if physPageSize == 0 {
		// The OS init code failed to fetch the physical page size.
		// 表示系统初始化失败，未能获取到物理页的大小
		throw("failed to get system page size")
	}
	// 检查物理页的大小是否合法
	if physPageSize < minPhysPageSize {
		// 检查和最小物理页的关系
		print("system page size (", physPageSize, ") is smaller than minimum page size (", minPhysPageSize, ")\n")
		throw("bad system page size")
	}
	if physPageSize&(physPageSize-1) != 0 {
		// 检查物理页和2的关系
		print("system page size (", physPageSize, ") must be a power of 2\n")
		throw("bad system page size")
	}

	// Initialize the heap.初始化堆，就是在初始化span，mcache等这些
	mheap_.init()
	// 获取主线程
	_g_ := getg()
	// 为主线程分配mcache
	_g_.m.mcache = allocmcache()

	// Create initial arena growth hints.
	if sys.PtrSize == 8 && GOARCH != "wasm" {
		// On a 64-bit machine, we pick the following hints
		// 在64为机器中我们使用这些东西
		// because:
		//
		// 1. Starting from the middle(中间) of the address space makes it easier to grow out a contiguous range without running in to some other mapping.
		//
		// 2. This makes Go heap addresses more easily recognizable when debugging.
		//
		// 3. Stack scanning in gccgo is still conservative,
		// so it's important that addresses be distinguishable from other data.
		//
		// Starting at 0x00c0 means that the valid(有效的) memory addresses
		// will begin 0x00c0, 0x00c1, ...
		// In little-endian, that's c0 00, c1 00, ... None of those are valid
		// UTF-8 sequences(序列), and they are otherwise as far away from ff (likely a common byte) as possible.
		// If that fails, we try other 0xXXc0 addresses.
		// An earlier attempt to use 0x11f8 caused out of memory errors on OS X during thread allocations.
		// 0x00c0 causes conflicts with AddressSanitizer which reserves all memory up to 0x0100.
		// These choices reduce the odds of a conservative garbage collector not collecting memory
		// because some non-pointer block of memory had a bit pattern that matched a memory address.
		// 这里应该说的是为什么size_class一半保存的是含有指针的，一半是不含有指针的
		// garbage collector 表示GC
		// However, on arm64, we ignore all this advice above and slam the allocation at 0x40 << 32
		// because when using 4k pages with 3-level translation buffers,
		// the user address space is limited to 39 bits On darwin/arm64, the address space is even smaller.
		for i := 0x7f; i >= 0; i-- {
			var p uintptr
			switch {
			case GOARCH == "arm64" && GOOS == "darwin":
				p = uintptr(i)<<40 | uintptrMask&(0x0013<<28)
			case GOARCH == "arm64":
				p = uintptr(i)<<40 | uintptrMask&(0x0040<<32)
			case raceenabled:
				// The TSAN runtime requires the heap to be in the range [0x00c000000000, 0x00e000000000).
				p = uintptr(i)<<32 | uintptrMask&(0x00c0<<32)
				if p >= uintptrMask&0x00e000000000 {
					continue
				}
			default:
				p = uintptr(i)<<40 | uintptrMask&(0x00c0<<32)
			}
			// 分配器,分配arena区域
			hint := (*arenaHint)(mheap_.arenaHintAlloc.alloc())
			hint.addr = p // 这个应该是起始地址
			hint.next, mheap_.arenaHints = mheap_.arenaHints, hint
		}
	} else {
		// On a 32-bit machine, we're much more concerned
		// about keeping the usable heap contiguous.
		// Hence:
		//
		// 1. We reserve space for all heapArenas up front so
		// they don't get interleaved with the heap. They're
		// ~258MB, so this isn't too bad. (We could reserve a
		// smaller amount of space up front if this is a
		// problem.)
		//
		// 2. We hint the heap to start right above the end of
		// the binary so we have the best chance of keeping it
		// contiguous.
		//
		// 3. We try to stake out a reasonably large initial
		// heap reservation.

		const arenaMetaSize = unsafe.Sizeof([1 << arenaBits]heapArena{})
		meta := uintptr(sysReserve(nil, arenaMetaSize))
		if meta != 0 {
			mheap_.heapArenaAlloc.init(meta, arenaMetaSize)
		}

		// We want to start the arena low, but if we're linked
		// against C code, it's possible global constructors
		// have called malloc and adjusted the process' brk.
		// Query the brk so we can avoid trying to map the
		// region over it (which will cause the kernel to put
		// the region somewhere else, likely at a high
		// address).
		procBrk := sbrk0()

		// If we ask for the end of the data segment but the
		// operating system requires a little more space
		// before we can start allocating, it will give out a
		// slightly higher pointer. Except QEMU, which is
		// buggy, as usual: it won't adjust the pointer
		// upward. So adjust it upward a little bit ourselves:
		// 1/4 MB to get away from the running binary image.
		p := firstmoduledata.end
		if p < procBrk {
			p = procBrk
		}
		if mheap_.heapArenaAlloc.next <= p && p < mheap_.heapArenaAlloc.end {
			p = mheap_.heapArenaAlloc.end
		}
		p = round(p+(256<<10), heapArenaBytes)
		// Because we're worried about fragmentation on
		// 32-bit, we try to make a large initial reservation.
		arenaSizes := []uintptr{
			512 << 20,
			256 << 20,
			128 << 20,
		}
		for _, arenaSize := range arenaSizes {
			a, size := sysReserveAligned(unsafe.Pointer(p), arenaSize, heapArenaBytes)
			if a != nil {
				mheap_.arena.init(uintptr(a), size)
				p = uintptr(a) + size // For hint below
				break
			}
		}
		hint := (*arenaHint)(mheap_.arenaHintAlloc.alloc())
		hint.addr = p
		hint.next, mheap_.arenaHints = mheap_.arenaHints, hint
	}
}

// sysAlloc allocates heap arena space for at least n bytes. 分配至少n个字节的堆空间
// The returned pointer is always heapArenaBytes-aligned and backed by h.arenas metadata.
// 返回值总是 _PageSize 对齐的(8k对齐)且处于arena_start和arena_end之间
// The returned size is always a multiple of heapArenaBytes.
// 返回的总是heapArenaBytes的倍数， heapArenaBytes是2的幂
// sysAlloc returns nil on failure.
// There is no corresponding free function.
// 没有对应的释放函数
// h must be locked.
func (h *mheap) sysAlloc(n uintptr) (v unsafe.Pointer, size uintptr) {
	n = round(n, heapArenaBytes)

	// First, try the arena pre-reservation. 尝试分配arena区域
	v = h.arena.alloc(n, heapArenaBytes, &memstats.heap_sys)
	if v != nil {
		size = n
		goto mapped
	}

	// Try to grow the heap at a hint address.
	// 这段代码的意思就是看看堆的可增长部分是否可以进行分配
	for h.arenaHints != nil {
		hint := h.arenaHints
		p := hint.addr
		if hint.down {
			// 表示arena可以扩展
			p -= n
		}
		if p+n < p {
			// We can't use this, so don't ask.
			v = nil
		} else if arenaIndex(p+n-1) >= 1<<arenaBits {
			// 表示取到的下标必须小于索引位， 不能越界
			// Outside addressable heap. Can't use.
			v = nil
		} else {
			v = sysReserve(unsafe.Pointer(p), n)
		}
		if p == uintptr(v) {
			// 这段表示分配成功，然后将hint的信息更新
			// Success. Update the hint.
			if !hint.down {
				p += n
			}
			hint.addr = p
			size = n
			break
		}
		// Failed. Discard this hint and try the next. 失败了会尝试下一次
		//
		// TODO: This would be cleaner if sysReserve could be told to only return the requested address.
		// In particular, this is already how Windows behaves, so it would simply things there.
		if v != nil {
			sysFree(v, n, nil)
		}
		h.arenaHints = hint.next
		h.arenaHintAlloc.free(unsafe.Pointer(hint))
	}
	// 再次检查size的大小
	if size == 0 {
		if raceenabled {
			// The race detector assumes the heap lives in
			// [0x00c000000000, 0x00e000000000), but we
			// just ran out of hints in this region. Give
			// a nice failure.
			throw("too many address space collisions for -race mode")
		}

		// All of the hints failed, so we'll take any (sufficiently aligned) address the kernel will give us.
		// 当所有的hits失败的时候我们将使用内核提供的任何地址(充分对齐)
		// 这块就是再从操作系统中获取内存
		v, size = sysReserveAligned(nil, n, heapArenaBytes)
		if v == nil {
			return nil, 0
		}

		// Create new hints for extending this region.
		// 创建一个新的hints为了扩展的区域
		hint := (*arenaHint)(h.arenaHintAlloc.alloc())
		hint.addr, hint.down = uintptr(v), true
		hint.next, mheap_.arenaHints = mheap_.arenaHints, hint
		hint = (*arenaHint)(h.arenaHintAlloc.alloc())
		hint.addr = uintptr(v) + size
		hint.next, mheap_.arenaHints = mheap_.arenaHints, hint
	}

	// Check for bad pointers or pointers we can't use.
	{
		var bad string
		p := uintptr(v)
		if p+size < p {
			bad = "region exceeds uintptr range"
		} else if arenaIndex(p) >= 1<<arenaBits {
			bad = "base outside usable address space"
		} else if arenaIndex(p+size-1) >= 1<<arenaBits {
			bad = "end outside usable address space"
		}
		if bad != "" {
			// This should be impossible on most architectures,
			// but it would be really confusing to debug.
			print("runtime: memory allocated by OS [", hex(p), ", ", hex(p+size), ") not in usable address space: ", bad, "\n")
			throw("memory reservation exceeds address space limit")
		}
	}

	if uintptr(v)&(heapArenaBytes-1) != 0 {
		throw("misrounded allocation in sysAlloc")
	}

	// Back the reservation.
	sysMap(v, size, &memstats.heap_sys)

mapped:
	// Create arena metadata.创建arena的元数据
	for ri := arenaIndex(uintptr(v)); ri <= arenaIndex(uintptr(v)+size-1); ri++ {
		l2 := h.arenas[ri.l1()]
		if l2 == nil {
			// Allocate an L2 arena map.给L2分配一个arena
			l2 = (*[1 << arenaL2Bits]*heapArena)(persistentalloc(unsafe.Sizeof(*l2), sys.PtrSize, nil))
			if l2 == nil {
				throw("out of memory allocating heap arena map")
			}
			// 利用原子性操作
			atomic.StorepNoWB(unsafe.Pointer(&h.arenas[ri.l1()]), unsafe.Pointer(l2))
		}

		if l2[ri.l2()] != nil {
			throw("arena already initialized")
		}
		var r *heapArena
		r = (*heapArena)(h.heapArenaAlloc.alloc(unsafe.Sizeof(*r), sys.PtrSize, &memstats.gc_sys))
		if r == nil {
			r = (*heapArena)(persistentalloc(unsafe.Sizeof(*r), sys.PtrSize, &memstats.gc_sys))
			if r == nil {
				throw("out of memory allocating heap arena metadata")
			}
		}

		// Store atomically just in case an object from the
		// new heap arena becomes visible before the heap lock
		// is released (which shouldn't happen, but there's little downside to this).
		// 这个设计就是防止出错，虽然基本不可能有问题
		atomic.StorepNoWB(unsafe.Pointer(&l2[ri.l2()]), unsafe.Pointer(r))
	}

	// Tell the race detector about the new heap memory.
	if raceenabled {
		racemapshadow(v, size)
	}

	return
}

// sysReserveAligned is like sysReserve, but the returned pointer is aligned to align bytes.
// It may reserve either n or n+align bytes, so it returns the size that was reserved.
// 返回的字节是对齐的
func sysReserveAligned(v unsafe.Pointer, size, align uintptr) (unsafe.Pointer, uintptr) {
	// Since the alignment is rather large in uses of this function, we're not likely to get it by chance,
	// so we ask for a larger region and remove the parts we don't need.
	retries := 0
retry:
	p := uintptr(sysReserve(v, size+align))
	switch {
	case p == 0:
		return nil, 0
	case p&(align-1) == 0:
		// We got lucky and got an aligned region, so we can
		// use the whole thing.
		return unsafe.Pointer(p), size + align
	case GOOS == "windows":
		// On Windows we can't release pieces of a
		// reservation, so we release the whole thing and
		// re-reserve the aligned sub-region. This may race,
		// so we may have to try again.
		sysFree(unsafe.Pointer(p), size+align, nil)
		p = round(p, align)
		p2 := sysReserve(unsafe.Pointer(p), size)
		if p != uintptr(p2) {
			// Must have raced. Try again.
			sysFree(p2, size, nil)
			if retries++; retries == 100 {
				throw("failed to allocate aligned heap memory; too many retries")
			}
			goto retry
		}
		// Success.
		return p2, size
	default:
		// Trim off the unaligned parts.
		pAligned := round(p, align)
		sysFree(unsafe.Pointer(p), pAligned-p, nil)
		end := pAligned + size
		endLen := (p + size + align) - end
		if endLen > 0 {
			sysFree(unsafe.Pointer(end), endLen, nil)
		}
		return unsafe.Pointer(pAligned), size
	}
}

// base address for all 0-byte allocations 所有0字节分配的基本地址
var zerobase uintptr

// nextFreeFast returns the next free object if one is quickly available.
// 返回下一个空闲可用对象（如果有一个快速可用的对象）
// Otherwise it returns 0.
func nextFreeFast(s *mspan) gclinkptr {
	// 查看第几位开始是0，即找到可用的
	theBit := sys.Ctz64(s.allocCache) // Is there a free object in the allocCache?
	// 如果超过了长度了则不进行处理，即返回0
	if theBit < 64 {
		result := s.freeindex + uintptr(theBit)
		if result < s.nelems {
			freeidx := result + 1
			if freeidx%64 == 0 && freeidx != s.nelems {
				return 0
			}
			s.allocCache >>= uint(theBit + 1)
			s.freeindex = freeidx
			s.allocCount++
			return gclinkptr(result*s.elemsize + s.base())
		}
	}
	return 0
}

// nextFree returns the next free object from the cached span if one is available.
// Otherwise it refills the cache with a span with an available object and
// returns that object along with a flag indicating that this was a heavy
// weight allocation. If it is a heavy weight allocation the caller must
// determine whether a new GC cycle needs to be started or if the GC is active
// whether this goroutine needs to assist the GC.
func (c *mcache) nextFree(spc spanClass) (v gclinkptr, s *mspan, shouldhelpgc bool) {
	s = c.alloc[spc]
	shouldhelpgc = false
	freeIndex := s.nextFreeIndex()
	if freeIndex == s.nelems {
		// The span is full.
		if uintptr(s.allocCount) != s.nelems {
			println("runtime: s.allocCount=", s.allocCount, "s.nelems=", s.nelems)
			throw("s.allocCount != s.nelems && freeIndex == s.nelems")
		}
		systemstack(func() {
			c.refill(spc)
		})
		shouldhelpgc = true
		s = c.alloc[spc]

		freeIndex = s.nextFreeIndex()
	}

	if freeIndex >= s.nelems {
		throw("freeIndex is not valid")
	}

	v = gclinkptr(freeIndex*s.elemsize + s.base())
	s.allocCount++
	if uintptr(s.allocCount) > s.nelems {
		println("s.allocCount=", s.allocCount, "s.nelems=", s.nelems)
		throw("s.allocCount > s.nelems")
	}
	return
}

// Allocate an object of size bytes.
// Small objects are allocated from the per-P cache's free lists. 小的objects通过p的cache分配
// Large objects (> 32 kB) are allocated straight from the heap.  大的objects直接在堆上分配
func mallocgc(size uintptr, typ *_type, needzero bool) unsafe.Pointer {
	if gcphase == _GCmarktermination {
		throw("mallocgc called with gcphase == _GCmarktermination")
	}

	if size == 0 {
		return unsafe.Pointer(&zerobase)
	}

	if debug.sbrk != 0 {
		align := uintptr(16)
		if typ != nil {
			align = uintptr(typ.align)
		}
		return persistentalloc(size, align, &memstats.other_sys)
	}

	// assistG is the G to charge for this allocation, or nil if
	// GC is not currently active.
	var assistG *g
	if gcBlackenEnabled != 0 {
		// Charge the current user G for this allocation.
		assistG = getg()
		if assistG.m.curg != nil {
			assistG = assistG.m.curg
		}
		// Charge the allocation against the G. We'll account
		// for internal fragmentation at the end of mallocgc.
		assistG.gcAssistBytes -= int64(size)

		if assistG.gcAssistBytes < 0 {
			// This G is in debt. Assist the GC to correct
			// this before allocating. This must happen
			// before disabling preemption.
			gcAssistAlloc(assistG)
		}
	}

	// Set mp.mallocing to keep from being preempted by GC.
	mp := acquirem()
	if mp.mallocing != 0 {
		throw("malloc deadlock")
	}
	if mp.gsignal == getg() {
		throw("malloc during signal")
	}
	mp.mallocing = 1

	shouldhelpgc := false
	dataSize := size
	c := gomcache()
	var x unsafe.Pointer
	noscan := typ == nil || typ.kind&kindNoPointers != 0
	if size <= maxSmallSize {
		if noscan && size < maxTinySize {
			// Tiny allocator.
			//
			// Tiny allocator combines several tiny allocation requests
			// into a single memory block. The resulting memory block
			// is freed when all subobjects are unreachable. The subobjects
			// must be noscan (don't have pointers), this ensures that
			// the amount of potentially wasted memory is bounded.
			//
			// Size of the memory block used for combining (maxTinySize) is tunable.
			// Current setting is 16 bytes, which relates to 2x worst case memory
			// wastage (when all but one subobjects are unreachable).
			// 8 bytes would result in no wastage at all, but provides less
			// opportunities for combining.
			// 32 bytes provides more opportunities for combining,
			// but can lead to 4x worst case wastage.
			// The best case winning is 8x regardless of block size.
			//
			// Objects obtained from tiny allocator must not be freed explicitly.
			// So when an object will be freed explicitly, we ensure that
			// its size >= maxTinySize.
			//
			// SetFinalizer has a special case for objects potentially coming
			// from tiny allocator, it such case it allows to set finalizers
			// for an inner byte of a memory block.
			//
			// The main targets of tiny allocator are small strings and
			// standalone escaping variables. On a json benchmark
			// the allocator reduces number of allocations by ~12% and
			// reduces heap size by ~20%.
			off := c.tinyoffset
			// Align tiny pointer for required (conservative) alignment.
			if size&7 == 0 {
				off = round(off, 8)
			} else if size&3 == 0 {
				off = round(off, 4)
			} else if size&1 == 0 {
				off = round(off, 2)
			}
			if off+size <= maxTinySize && c.tiny != 0 {
				// The object fits into existing tiny block.
				x = unsafe.Pointer(c.tiny + off)
				c.tinyoffset = off + size
				c.local_tinyallocs++
				mp.mallocing = 0
				releasem(mp)
				return x
			}
			// Allocate a new maxTinySize block.
			span := c.alloc[tinySpanClass]
			v := nextFreeFast(span)
			if v == 0 {
				v, _, shouldhelpgc = c.nextFree(tinySpanClass)
			}
			x = unsafe.Pointer(v)
			(*[2]uint64)(x)[0] = 0
			(*[2]uint64)(x)[1] = 0
			// See if we need to replace the existing tiny block with the new one
			// based on amount of remaining free space.
			if size < c.tinyoffset || c.tiny == 0 {
				c.tiny = uintptr(x)
				c.tinyoffset = size
			}
			size = maxTinySize
		} else {
			var sizeclass uint8
			if size <= smallSizeMax-8 {
				sizeclass = size_to_class8[(size+smallSizeDiv-1)/smallSizeDiv]
			} else {
				sizeclass = size_to_class128[(size-smallSizeMax+largeSizeDiv-1)/largeSizeDiv]
			}
			size = uintptr(class_to_size[sizeclass])
			spc := makeSpanClass(sizeclass, noscan)
			span := c.alloc[spc]
			v := nextFreeFast(span)
			if v == 0 {
				v, span, shouldhelpgc = c.nextFree(spc)
			}
			x = unsafe.Pointer(v)
			if needzero && span.needzero != 0 {
				memclrNoHeapPointers(unsafe.Pointer(v), size)
			}
		}
	} else {
		var s *mspan
		shouldhelpgc = true
		systemstack(func() {
			s = largeAlloc(size, needzero, noscan)
		})
		s.freeindex = 1
		s.allocCount = 1
		x = unsafe.Pointer(s.base())
		size = s.elemsize
	}

	var scanSize uintptr
	if !noscan {
		// If allocating a defer+arg block, now that we've picked a malloc size
		// large enough to hold everything, cut the "asked for" size down to
		// just the defer header, so that the GC bitmap will record the arg block
		// as containing nothing at all (as if it were unused space at the end of
		// a malloc block caused by size rounding).
		// The defer arg areas are scanned as part of scanstack.
		if typ == deferType {
			dataSize = unsafe.Sizeof(_defer{})
		}
		heapBitsSetType(uintptr(x), size, dataSize, typ)
		if dataSize > typ.size {
			// Array allocation. If there are any
			// pointers, GC has to scan to the last
			// element.
			if typ.ptrdata != 0 {
				scanSize = dataSize - typ.size + typ.ptrdata
			}
		} else {
			scanSize = typ.ptrdata
		}
		c.local_scan += scanSize
	}

	// Ensure that the stores above that initialize x to
	// type-safe memory and set the heap bits occur before
	// the caller can make x observable to the garbage
	// collector. Otherwise, on weakly ordered machines,
	// the garbage collector could follow a pointer to x,
	// but see uninitialized memory or stale heap bits.
	publicationBarrier()

	// Allocate black during GC.
	// All slots hold nil so no scanning is needed.
	// This may be racing with GC so do it atomically if there can be
	// a race marking the bit.
	if gcphase != _GCoff {
		gcmarknewobject(uintptr(x), size, scanSize)
	}

	if raceenabled {
		racemalloc(x, size)
	}

	if msanenabled {
		msanmalloc(x, size)
	}

	mp.mallocing = 0
	releasem(mp)

	if debug.allocfreetrace != 0 {
		tracealloc(x, size, typ)
	}

	if rate := MemProfileRate; rate > 0 {
		if size < uintptr(rate) && int32(size) < c.next_sample {
			c.next_sample -= int32(size)
		} else {
			mp := acquirem()
			profilealloc(mp, x, size)
			releasem(mp)
		}
	}

	if assistG != nil {
		// Account for internal fragmentation in the assist
		// debt now that we know it.
		assistG.gcAssistBytes -= int64(size - dataSize)
	}

	if shouldhelpgc {
		if t := (gcTrigger{kind: gcTriggerHeap}); t.test() {
			gcStart(gcBackgroundMode, t)
		}
	}

	return x
}

func largeAlloc(size uintptr, needzero bool, noscan bool) *mspan {
	// print("largeAlloc size=", size, "\n")

	if size+_PageSize < size {
		throw("out of memory")
	}
	npages := size >> _PageShift
	if size&_PageMask != 0 {
		npages++
	}

	// Deduct credit for this span allocation and sweep if
	// necessary. mHeap_Alloc will also sweep npages, so this only
	// pays the debt down to npage pages.
	deductSweepCredit(npages*_PageSize, npages)

	s := mheap_.alloc(npages, makeSpanClass(0, noscan), true, needzero)
	if s == nil {
		throw("out of memory")
	}
	s.limit = s.base() + size
	heapBitsForAddr(s.base()).initSpan(s)
	return s
}

// implementation of new builtin
// compiler (both frontend and SSA backend) knows the signature
// of this function
func newobject(typ *_type) unsafe.Pointer {
	return mallocgc(typ.size, typ, true)
}

//go:linkname reflect_unsafe_New reflect.unsafe_New
func reflect_unsafe_New(typ *_type) unsafe.Pointer {
	return mallocgc(typ.size, typ, true)
}

// newarray allocates an array of n elements of type typ.
func newarray(typ *_type, n int) unsafe.Pointer {
	if n == 1 {
		return mallocgc(typ.size, typ, true)
	}
	if n < 0 || uintptr(n) > maxSliceCap(typ.size) {
		panic(plainError("runtime: allocation size out of range"))
	}
	return mallocgc(typ.size*uintptr(n), typ, true)
}

//go:linkname reflect_unsafe_NewArray reflect.unsafe_NewArray
func reflect_unsafe_NewArray(typ *_type, n int) unsafe.Pointer {
	return newarray(typ, n)
}

func profilealloc(mp *m, x unsafe.Pointer, size uintptr) {
	mp.mcache.next_sample = nextSample()
	mProf_Malloc(x, size)
}

// nextSample returns the next sampling point for heap profiling. The goal is
// to sample allocations on average every MemProfileRate bytes, but with a
// completely random distribution over the allocation timeline; this
// corresponds to a Poisson process with parameter MemProfileRate. In Poisson
// processes, the distance between two samples follows the exponential
// distribution (exp(MemProfileRate)), so the best return value is a random
// number taken from an exponential distribution whose mean is MemProfileRate.
func nextSample() int32 {
	if GOOS == "plan9" {
		// Plan 9 doesn't support floating point in note handler.
		if g := getg(); g == g.m.gsignal {
			return nextSampleNoFP()
		}
	}

	return fastexprand(MemProfileRate)
}

// fastexprand returns a random number from an exponential distribution with
// the specified mean.
func fastexprand(mean int) int32 {
	// Avoid overflow. Maximum possible step is
	// -ln(1/(1<<randomBitCount)) * mean, approximately 20 * mean.
	switch {
	case mean > 0x7000000:
		mean = 0x7000000
	case mean == 0:
		return 0
	}

	// Take a random sample of the exponential distribution exp(-mean*x).
	// The probability distribution function is mean*exp(-mean*x), so the CDF is
	// p = 1 - exp(-mean*x), so
	// q = 1 - p == exp(-mean*x)
	// log_e(q) = -mean*x
	// -log_e(q)/mean = x
	// x = -log_e(q) * mean
	// x = log_2(q) * (-log_e(2)) * mean    ; Using log_2 for efficiency
	const randomBitCount = 26
	q := fastrand()%(1<<randomBitCount) + 1
	qlog := fastlog2(float64(q)) - randomBitCount
	if qlog > 0 {
		qlog = 0
	}
	const minusLog2 = -0.6931471805599453 // -ln(2)
	return int32(qlog*(minusLog2*float64(mean))) + 1
}

// nextSampleNoFP is similar to nextSample, but uses older,
// simpler code to avoid floating point.
func nextSampleNoFP() int32 {
	// Set first allocation sample size.
	rate := MemProfileRate
	if rate > 0x3fffffff { // make 2*rate not overflow
		rate = 0x3fffffff
	}
	if rate != 0 {
		return int32(fastrand() % uint32(2*rate))
	}
	return 0
}

type persistentAlloc struct {
	base *notInHeap
	off  uintptr
}

var globalAlloc struct {
	mutex
	persistentAlloc
}

// Wrapper around sysAlloc that can allocate small chunks.
// There is no associated free operation.
// Intended for things like function/type/debug-related persistent data.
// If align is 0, uses default align (currently 8).
// The returned memory will be zeroed.
//
// Consider marking persistentalloc'd types go:notinheap.
func persistentalloc(size, align uintptr, sysStat *uint64) unsafe.Pointer {
	var p *notInHeap
	systemstack(func() {
		p = persistentalloc1(size, align, sysStat)
	})
	return unsafe.Pointer(p)
}

// Must run on system stack because stack growth can (re)invoke it.
// See issue 9174.
//go:systemstack
func persistentalloc1(size, align uintptr, sysStat *uint64) *notInHeap {
	const (
		chunk    = 256 << 10
		maxBlock = 64 << 10 // VM reservation granularity is 64K on windows
	)

	if size == 0 {
		throw("persistentalloc: size == 0")
	}
	if align != 0 {
		if align&(align-1) != 0 {
			throw("persistentalloc: align is not a power of 2")
		}
		if align > _PageSize {
			throw("persistentalloc: align is too large")
		}
	} else {
		align = 8
	}

	if size >= maxBlock {
		return (*notInHeap)(sysAlloc(size, sysStat))
	}

	mp := acquirem()
	var persistent *persistentAlloc
	if mp != nil && mp.p != 0 {
		persistent = &mp.p.ptr().palloc
	} else {
		lock(&globalAlloc.mutex)
		persistent = &globalAlloc.persistentAlloc
	}
	persistent.off = round(persistent.off, align)
	if persistent.off+size > chunk || persistent.base == nil {
		persistent.base = (*notInHeap)(sysAlloc(chunk, &memstats.other_sys))
		if persistent.base == nil {
			if persistent == &globalAlloc.persistentAlloc {
				unlock(&globalAlloc.mutex)
			}
			throw("runtime: cannot allocate memory")
		}
		persistent.off = 0
	}
	p := persistent.base.add(persistent.off)
	persistent.off += size
	releasem(mp)
	if persistent == &globalAlloc.persistentAlloc {
		unlock(&globalAlloc.mutex)
	}

	if sysStat != &memstats.other_sys {
		mSysStatInc(sysStat, size)
		mSysStatDec(&memstats.other_sys, size)
	}
	return p
}

// linearAlloc is a simple linear allocator that pre-reserves a region
// of memory and then maps that region as needed. The caller is
// responsible for locking.
type linearAlloc struct {
	next   uintptr // next free byte
	mapped uintptr // one byte past end of mapped space
	end    uintptr // end of reserved space
}

func (l *linearAlloc) init(base, size uintptr) {
	l.next, l.mapped = base, base
	l.end = base + size
}

func (l *linearAlloc) alloc(size, align uintptr, sysStat *uint64) unsafe.Pointer {
	p := round(l.next, align)
	if p+size > l.end {
		return nil
	}
	l.next = p + size
	if pEnd := round(l.next-1, physPageSize); pEnd > l.mapped {
		// We need to map more of the reserved(预留) space.
		sysMap(unsafe.Pointer(l.mapped), pEnd-l.mapped, sysStat)
		l.mapped = pEnd
	}
	return unsafe.Pointer(p)
}

// notInHeap is off-heap memory allocated by a lower-level allocator
// like sysAlloc or persistentAlloc.
//
// In general, it's better to use real types marked as go:notinheap,
// but this serves as a generic type for situations where that isn't
// possible (like in the allocators).
//
// TODO: Use this as the return type of sysAlloc, persistentAlloc, etc?
//
//go:notinheap
type notInHeap struct{}

func (p *notInHeap) add(bytes uintptr) *notInHeap {
	return (*notInHeap)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + bytes))
}
