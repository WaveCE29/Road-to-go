package main

import "fmt"

type employee struct {
	employeeID   string
	employeeName string
	phone        string
}

func main() {

	//array of struct
	// employeeList := [3]employee{}
	// employeeList[0] = employee{
	// 	employeeID:   "E001",
	// 	employeeName: "John Doe",
	// 	phone:        "1234567890"}
	// employeeList[1] = employee{
	// 	employeeID:   "E002",
	// 	employeeName: "Jane Doe",
	// 	phone:        "0987654321"}
	// employeeList[2] = employee{
	// 	employeeID:   "E003",
	// 	employeeName: "John Smith",
	// 	phone:        "1230984567"}
	// fmt.Println("Employee List = ", employeeList)

	//slice of struct
	employeeList := []employee{}
	employee1 := employee{
		employeeID:   "E001",
		employeeName: "John Doe",
		phone:        "1234567890"}
	employee2 := employee{
		employeeID:   "E002",
		employeeName: "Jane Doe",
		phone:        "0987654321"}
	employee3 := employee{
		employeeID:   "E003",
		employeeName: "John Smith",
		phone:        "1230984567"}

	employeeList = append(employeeList, employee1)
	employeeList = append(employeeList, employee2)
	employeeList = append(employeeList, employee3)
	fmt.Println("Employee List = ", employeeList)
}
