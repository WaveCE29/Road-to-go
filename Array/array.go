package main

import "fmt"

var products [4]string

//var price [4]float64

func main() {
	products[0] = "Computer"
	products[1] = "Mouse"
	products[2] = "Keyboard"
	products[3] = "Monitor"

	price := [4]float64{1000, 50, 30, 500}

	fmt.Println(products)
	fmt.Println(price)

}
