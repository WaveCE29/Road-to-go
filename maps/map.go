package main

import "fmt"

var products = make(map[string]float64)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Products:", products)
	products["Computer"] = 1000
	products["Mouse"] = 50
	products["Keyboard"] = 30
	products["Monitor"] = 500
	fmt.Println("Products:", products)

	// delete a key
	delete(products, "Mouse")
	fmt.Println("Products:", products)

	// update a key
	products["Computer"] = 2000
	fmt.Println("Products:", products)

	// check if a key exists
	val1 := products["Computer"]
	fmt.Println("Computer:", val1)

	courseName := map[string]string{"Python": "Intermediate", "Java": "Advanced"}
	fmt.Println("Course Name:", courseName)

}
