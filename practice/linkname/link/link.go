package link

import _"unsafe"

//go:linkname helloWorld readruntime/practice/linkname/hello/Hello
func helloWorld(){
	println("Hello world!")
}
