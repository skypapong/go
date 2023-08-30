package main

import "fmt"

var poin int

func main() {

	fmt.Println("กรอกคะแนน")
	fmt.Scanf("%d", &poin)
	fmt.Println("คะแนนที่ได้:", poin)

	if poin >= 80 {
		fmt.Println("Grad A")
	} else if poin >= 70 {
		if poin > 75 {
			fmt.Println("Grad B+")
		} else {
			fmt.Println("Grad B")
		}
	} else if poin >= 60 {
		if poin > 65 {
			fmt.Println("Grad C+")
		} else {
			fmt.Println("Grad C")
		}
	} else if poin >= 50 {
		if poin > 55 {
			fmt.Println("Grad D+")
		} else {
			fmt.Println("Grad D")
		}
	} else {
		fmt.Println("Grad F")
	}
}
