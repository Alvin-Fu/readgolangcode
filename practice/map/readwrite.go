package main

import "sync"

// map是可以并发的进行读，但是不能并发的写，当然并发读写也是不行的
func main() {
	var rw = make(map[int]int)
	wg := sync.WaitGroup{}
	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go func(i int) {
			rw[i] = i + 1
			wg.Done()
		}(i)
		go func(i int) {
			_ = rw[i]
			wg.Done()
		}(i)
	}
	wg.Wait()
}
