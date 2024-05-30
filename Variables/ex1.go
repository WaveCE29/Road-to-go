package main

import "fmt"

func main() {

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	numFloat := 3.14
	fmt.Println(numFloat)

	fmt.Println(float64(c) + numFloat)

	fmt.Println("Hello, World!", numFloat)
}
