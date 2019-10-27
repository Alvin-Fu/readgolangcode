package main

func Add()int{
	num := 0
	for i := 0; i < 10; i++{
		num += i
	}
	return num
}

func Sum()*int{
	num := 100
	for i := 0; i < 10; i++{
		num -= i
	}
	return &num
}

func main(){
	Add()
	Sum()
}

