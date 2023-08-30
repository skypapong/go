package main

import (
	"fmt"
)

func main() {
	var coseName []string
	coseName = []string{"JAVA", "GO"}
	fmt.Println(coseName)
	coseName = append(coseName, "C", "C++", "C#", "HTML", "PYTHON")
	fmt.Println(coseName)

	coseWeb := coseName[4:7]
	fmt.Println(coseWeb)

	coseWeb = coseName[:4]
	fmt.Println(coseWeb)
}
