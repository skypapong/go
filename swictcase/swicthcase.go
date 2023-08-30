package main

import "fmt"

func main() {
	var input int
	// input := 7
	fmt.Print("กรอกค่า")
	fmt.Scanf("%d", &input)
	fmt.Println("input", input)

	switch input {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("NO")
	}
}
