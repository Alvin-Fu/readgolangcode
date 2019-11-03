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

	chan 在使用range的时候这个chan必须是被关闭的， 否则会产生死锁
*/

func main() {
	ChanRange()
}

func ChanWriteRead() {
	ch := make(chan int, 2)
	//go func() {
	ch <- 42
	//}()
	va := <-ch
	fmt.Println(va)
}

// chan 在使用range的时候这个chan必须是被关闭的
func ChanRange() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
	}()
	for v := range ch {
		fmt.Println(v)
	}
}

//参考的博客
// https://guzalexander.com/2013/12/06/golang-channels-tutorial.html
