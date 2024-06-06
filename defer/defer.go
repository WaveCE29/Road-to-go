package main

import "fmt"

// func add(value1, value2 int) {
// 	result := value1 + value2
// 	fmt.Println("Addition of", value1, "and", value2, "is", result)

// }

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Println("Loop i :", i)
	}
}

func deferloop() {
	for j := 0; j < 10; j++ {
		defer fmt.Println("Loop j :", j)
	}
}

// defer work in LIFO order

func main() {
	// fmt.Println("Start")
	// defer fmt.Println("End")
	// defer add(10, 20)
	// defer add(30, 40)
	// fmt.Println("Middle")
	loop()
	deferloop()
}
