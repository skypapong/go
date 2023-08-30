package main

import "fmt"

//นิยามตัวแปร

func main2() {
	var score int
	fmt.Printf("ป้อนคะแนน = ")
	fmt.Scanf("%d", &score)

	fmt.Println("คะแนน", score)

	if score == 1 {
		fmt.Println("ผ่าน")
	} else if score == 2 {
		fmt.Println("ไม่ผ่่าน")
	} else {
		fmt.Println("โกง")
	}

	switch score {
	case 1:
		fmt.Println("ผ่าน")
	case 2:
		fmt.Println("ไม่ผ่่าน")
	default:
		fmt.Println("โกง")
	}

	num1, num2 := 10, 3
	fmt.Println("เท่ากันหรือไม่", num1 == num2)
	fmt.Println(num1, "มากกว่า", num2, "=", num1 > num2)
}
