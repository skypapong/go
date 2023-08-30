package main

import "fmt"

type employee struct {
	employeeID   string
	employeeName string
	Phone        string
}

func main() {

	// employee1 := employee{
	// 	employeeID:   "101",
	// 	employeeName: "Sky",
	// 	Phone:        "0988888888",
	// }
	// fmt.Println("employee1 = ", employee1)

	// employeeList := [3]employee{}
	// employeeList[0] = employee{
	// 	employeeID:   "101",
	// 	employeeName: "Sky",
	// 	Phone:        "0988888888",
	// }
	// employeeList[1] = employee{
	// 	employeeID:   "102",
	// 	employeeName: "Din",
	// 	Phone:        "0988888888",
	// }
	// employeeList[2] = employee{
	// 	employeeID:   "103",
	// 	employeeName: "Oom",
	// 	Phone:        "0988888888",
	// }
	// fmt.Println("employee = ", employeeList)

	employeeList := []employee{}
	employee1 := employee{
		employeeID:   "101",
		employeeName: "fon",
		Phone:        "0655328891",
	}
	employee2 := employee{
		employeeID:   "102",
		employeeName: "boom",
		Phone:        "0663113131",
	}
	employee3 := employee{
		employeeID:   "102",
		employeeName: "dom",
		Phone:        "0989808800",
	}

	employeeList = append(employeeList, employee1)
	employeeList = append(employeeList, employee2)
	employeeList = append(employeeList, employee3)

	fmt.Println("employee = ", employeeList)

}
