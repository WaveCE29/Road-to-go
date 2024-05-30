package main

import "fmt"

func hello() {
	fmt.Println("Hello World!")
}

func plus(val1 int, val2 int) {
	result := val1 + val2
	fmt.Println("Result:", result)
}

func plus2(val1 int, val2 int) int {
	result := val1 + val2
	return result
}

func plus3() int {
	var val1 int = 10
	var val2 int = 30
	result := val1 + val2
	return result
}

func main() {
	hello()
	plus(10, 20)
	X := plus2(10, 20)
	fmt.Println("X:", X)
	Y := plus3()
	fmt.Println("Y:", Y)
}
