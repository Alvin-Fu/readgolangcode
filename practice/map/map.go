package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	FloatMap()
}
func FloatMap() {
	fMap := make(map[float64]int)
	rand.Seed(time.Now().Unix())
	fArr := make([]float64, 0)
	count := 10
	for i := 0; i < count; i++ {
		fArr = append(fArr, rand.Float64())
		//time.Sleep(1 * time.Second)
	}
	for i := 0; i < count; i++ {
		fMap[fArr[i]] = i
		fmt.Println(fMap[fArr[i]], fArr[i])
	}
	for i := 0; i < count; i++ {
		fmt.Println(fMap[fArr[i]], fArr[i])
	}
}
