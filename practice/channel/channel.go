package main

import "fmt"

//使用channel的时候注意的点：
/*
 *写操作：对于一个无缓存的channel，写操作能够成功必须有一个读的goroutine在等待，因为写操作是阻塞的，必须等到读才能继续运行
 *   ch := make(chan int, 2)
	//go func() {
	ch <- 42
	//}()
	va := <-ch
	fmt.Println(va)
	会产生死锁
*/

func main() {
	ch := make(chan int, 2)
	//go func() {
	ch <- 42
	//}()
	va := <-ch
	fmt.Println(va)
}
