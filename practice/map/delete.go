package main

import (
	"fmt"
	"math"
)

func main() {
	demo2()
}

/*func demo1() {
	var dm = make(map[int]int)
	for i := 1; i < 11; i++ {
		dm[i] = i
	}
	wg := sync.WaitGroup{}
	wg.Add(5)
	for k, _ := range dm {
		if k == 5 {

			go func() {
				for i := 12; i < 21; i++ {
					dm[i] = i
				}
				dm[11] = 11
				for i := 12; i < 16; i++ {
					go func() {
						delete(dm, i)
						wg.Done()
					}()
				}
				wg.Done()
			}()
		}
		fmt.Println(k)
	}
	wg.Wait()
	fmt.Println(len(dm), dm)
}*/

//在遍历的过程中在做插入的时候，新插入的有可能会被遍历到，也有可能不能被遍历到，这个要取决于，新插入所在的桶是否已经被遍历
// 在遍历的时候进行删除也会有这些问题
func demo2() {
	var dm = make(map[int]int)
	for i := 1; i < 11; i++ {
		dm[i] = i
	}
	//wg := sync.WaitGroup{}
	//wg.Add(5)
	num := 0
	for k, _ := range dm {
		if k == 5 {
			for i := 12; i < 21; i++ {
				dm[i] = i
			}
		}
		num++
		fmt.Println(k)
	}
	fmt.Println(num)
	//wg.Wait()
	fmt.Println(len(dm), dm)

}
