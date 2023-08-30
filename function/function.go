package main

import "fmt"

func hello() {
	fmt.Println("Hello Dev")
}

func plus(num1 int, num2 int) {
	resule := num1 + num2
	fmt.Println("resule ==", resule)
}

func plus2(num1 int, num2 int) int {
	return num1 + num2
}

func plus3(num1, num2, num3 int) int {
	return num1 + num2 + num3
}

func main() {
	hello()
	plus(10, 20)
	total := plus2(1, 2)
	fmt.Println("total ==", total)
	total2 := plus3(1, 2, 3)
	fmt.Println("total2 ==", total2)
}
