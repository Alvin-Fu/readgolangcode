package main

import "fmt"

func main() {
	var t = 0
	var p *int

	p = &t
	fmt.Println(p)
	fmt.Println(uint8(1))

	fmt.Println(3 &^ 5)
	//for i := 0; i < 10; i++ {
	//	for j := 0; j < 10; j++ {
	//		continue
	//	}
	//	fmt.Println(i)
	//}
	//fmt.Println("hello")
}
