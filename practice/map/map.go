package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	IsFloat()
	//FloatMap()
}
func IsFloat() {
	rand.Seed(time.Now().Unix())
	fmt.Println(math.Float64bits(rand.Float64()))
	if math.Float64bits(math.Inf(100)) == math.Float64bits(math.Inf(1000)) {
		fmt.Println("hello")
	}
	fmt.Printf("%v\n", math.Float64bits(math.Inf(89989)))
	fmt.Printf("%v\n", math.Float64bits(math.Inf(89)))
	fmt.Printf("%v\n", math.NaN())
}

func FloatMap() {
	fMap := make(map[float64]int)
	rand.Seed(time.Now().Unix())
	fArr := make([]float64, 0)
	count := 10
	fmt.Println(math.IsNaN(rand.Float64()))
	for i := 0; i < count; i++ {
		fArr = append(fArr, math.Inf(i))
		//time.Sleep(1 * time.Second)
	}
	fmt.Printf("%+v\n", fArr)
	for i := 0; i < count; i++ {
		fMap[fArr[i]] = i
		fmt.Println(fMap[fArr[i]], fArr[i])
	}
	for i := 0; i < count; i++ {
		fmt.Println(fMap[fArr[i]], fArr[i])
	}
}
