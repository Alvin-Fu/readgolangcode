package main

import (
	"fmt"
	"reflect"
)

func main() {
	var r1 interface{}
	var r2 interface{}
	var t int64 = 1
	r1 = 1
	r2 = t
	fmt.Println(reflect.DeepEqual(r1, r2))

}
