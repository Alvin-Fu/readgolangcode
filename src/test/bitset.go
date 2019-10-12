package main

import (
	"fmt"
)

const (
	shift = 6
	mask  = 0x3f // 即0b00111111
)

type Bitset struct {
	data []int64
}

func NewBitSet(n int) *Bitset {
	// 获取位置信息
	index := n >> shift
	fmt.Println(index)
	set := &Bitset{
		data: make([]int64, index+1),
	}
	fmt.Println(1 << uint(n&mask))
	// 根据 n 设置 bitset
	set.data[index] |= 1 << uint(n&mask)
	fmt.Println(set.data)
	return set
}

func (set *Bitset) Contains(n int) bool {
	// 获取位置信息
	index := n >> shift
	fmt.Println(1<<uint(n&mask), n&mask)
	return set.data[index]&(1<<uint(n&mask)) != 0
}

func main() {
	set := NewBitSet(10)
	fmt.Println("set contains 65", set.Contains(10))
	fmt.Println("set contains 64", set.Contains(64))
}
