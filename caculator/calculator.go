package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getInput(promt string) float64 {
	fmt.Printf("%v", promt)
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		message, _ := fmt.Scanf("%v must number only", promt)
		panic(message)
	}
	return value
}

func add(value1, value2 float64) float64 {
	return value1 + value2
}

func delete(value1, value2 float64) float64 {
	return value1 - value2
}

func cal(value1, value2 float64) float64 {
	return value1 * value2
}

func mod(value1, value2 float64) float64 {
	return value1 / value2
}

func getOperator() string {
	fmt.Println("operator is (+ - * /):")
	op, _ := reader.ReadString('\n')
	return strings.TrimSpace(op)
}

func main() {
	value1 := getInput("value1 =")
	value2 := getInput("value2 =")

	var result float64
	switch operator := getOperator(); operator {
	case "+":
		result = add(value1, value2)
	case "-":
		result = delete(value1, value2)
	case "*":
		result = cal(value1, value2)
	case "/":
		result = mod(value1, value2)
	default:
		panic("wrong operator")
	}
	fmt.Printf("result = %v", result)
}
