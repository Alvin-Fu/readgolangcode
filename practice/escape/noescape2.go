package main

import "fmt"

type Tmp struct {
	a, b int
}
const num1, num2 = 100, 200
//go:inline
func Cen(t *Tmp){
	t.a = num1 + num2
	t.b = num2 - num1
}

func main(){
	tmp := new(Tmp)
	Cen(tmp)
	fmt.Println(tmp)
}
