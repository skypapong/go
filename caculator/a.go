package main

import "fmt"

func main() {
	fmt.Println("couting")

	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
	// 1700 1900 1900 1500 2700  1800 800 800 3700
}
