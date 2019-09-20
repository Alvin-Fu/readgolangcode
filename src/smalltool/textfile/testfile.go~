package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

func main() {
	file, err := os.Create("text.txt")
	if err != nil {
		fmt.Println("File was created fail: %v", err)
		return
	}
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(file *os.File, i int) {
			str := strconv.Itoa(i)
			for j := 0; j < 5; j++ {
				str += str
			}
			_, err := file.Write([]byte(str))
			if err != nil {
				fmt.Println(err)
			}
			file.Write([]byte("-----" + str + "h\n"))
			wg.Done()
		}(file, i)
	}
	wg.Wait()
}
