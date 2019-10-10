package main

import "fmt"

func main() {
	n := 99
	a := 4
	fmt.Println(n + a - 1)
	fmt.Println(a - 1)
	fmt.Println((n + a - 1) &^ (a - 1))
}
