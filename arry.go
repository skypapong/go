package main

import (
	"fmt"
)

// struct
type Product struct {
	name     string
	price    float64
	category string
	discount int
}

func main() {
	// var number [3]int = [3]int{100, 200, 300}
	// number := []int{100, 200, 300, 400, 500, 600, 700, 800, 900}

	// for i := 0; i < len(number); i++ {
	// 	fmt.Println("number int =", i, "value", number[i])
	// 	if i == 4 {
	// 		break
	// 	}
	// }

	// for index, value := range number {
	// 	fmt.Println("number int =", index, "value", value)
	// }

	ShowText("Sky")
	DD(10, 10)
	a := AA()
	fmt.Println("AA=", a)
	total, check := CC(200, 100)
	fmt.Println("Total ==", total, check)
	result := Varidic(10, 20, 30)
	fmt.Println("Variadic", result)

}

// var a int = 5
// var b int = 10

func DD(a int, b int) {

	c := a + b
	fmt.Println("all==", c)

}

func CC(a int, b int) (int, string) {

	c := a + b
	status := ""
	if c%2 == 0 {
		status = "เลขคู่"
	} else {
		status = "เลขคี่"
	}
	return c, status

}

func Varidic(number ...int) int {
	total := 0
	for _, value := range number {
		total += value
	}
	return total
}

func ShowText(name string) {
	fmt.Println("Hello:", name)
}

func AA() int {
	return 20
}
