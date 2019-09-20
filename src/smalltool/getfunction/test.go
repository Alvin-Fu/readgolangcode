package main

import "fmt"
import (
	"runtime"
	"runtime/debug"
)

func main() {
	pc, _, _, _ := runtime.Caller(0)
	fmt.Println("Name of function: " + runtime.FuncForPC(pc).Name())
	fmt.Println()

	// or, define a function for it
	fmt.Println("Name of function: " + funcName())
	//d()
	//x()
	test1()
}

func funcName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

func x() {
	fmt.Println("Name of function: " + funcName())
	y()
	d()
}

func y() {
	fmt.Println("Name of function: " + funcName())
	z()
}
func z() {
	fmt.Println("Name of function: " + funcName())
}
func d() {
	fmt.Println("Name of function: " + funcName())
}
func test1() {
	test2()
}

func test2() {
	test3()
}

func test3() {
	fmt.Printf("%s", debug.Stack())
	debug.PrintStack()
}
