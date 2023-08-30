package main

import "fmt"

var product = make(map[string]float32)

func main() {
	fmt.Println("product==", product)

	//add

	product["Mackbook"] = 40000
	product["Ipone"] = 20000
	product["Ipad"] = 30000
	fmt.Println("product==", product)

	//delete
	delete(product, "Ipone")
	fmt.Println("product==", product)

	//update
	product["Ipad"] = 20000
	fmt.Println("product==", product)

	value1 := product["Ipad"]
	fmt.Println("value1==", value1)

	couseName := map[string]string{"101": "JAVA", "102": "Python"}
	fmt.Println("couseName==", couseName)

}
