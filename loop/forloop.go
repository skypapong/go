package main

import "fmt"

func main() {
	// i := 0
	// for i < 10 {
	// 	fmt.Println("i=", i)
	// 	i = i + 1
	// }

	// for a := 0; a < 15; a++ {
	// 	fmt.Println("a=", a)
	// }

	var input string

	for {
		fmt.Scanf("%s", &input)
		fmt.Println("input=", input)
		if input == "exit" {
			break
		}
	}
}
