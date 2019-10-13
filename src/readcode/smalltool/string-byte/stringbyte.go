package main

import (
	"bytes"
	"fmt"
)

func main() {
	str := "he0llo0"
	var by [100]byte
	for i := 0; i < 100; i++ {
		by[i] = 0
	}
	fmt.Println(by)
	b := []byte(str)
	for i, v := range b {
		by[i] = v
	}
	fmt.Println(bytes.IndexByte(by[:], 0))
	fmt.Println(string(by[:]))
}
