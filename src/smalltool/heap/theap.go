package main

import "fmt"

const PtrSize = 4 << (^uintptr(0) >> 63)

func main() {
	n := 64 * 1024 / PtrSize
	fmt.Println(n)
}
