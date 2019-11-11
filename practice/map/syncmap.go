package main

import (
	"fmt"

	"sync"
	"unsafe"
)

func main() {
	var tmp = sync.Map{}
	tmp.Store(5, 10)

	//tmp.Range(func(key, value interface{}) bool {
	//	fmt.Println(key, value)
	//	return true
	//})
	tmp.Delete(5)
	fmt.Println(tmp.LoadOrStore(4, 10))
	fmt.Println(tmp.LoadOrStore(5, 10))
	tmp.Range(func(key, value interface{}) bool {
		fmt.Println("second: ", key, value)
		return true
	})
	//tmp.Delete(5)
	tmp.Delete(4)
	tmp.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
	var expunged = unsafe.Pointer(new(interface{}))
	t := unsafe.Pointer(new(interface{}))
	if expunged == t {
		fmt.Println("hello")
	}
	fmt.Println(expunged, unsafe.Pointer(new(interface{})))
}
