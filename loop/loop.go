package main

import "fmt"

func main() {
	// i := 0
	// for i < 5 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// for j := 0; j < 5; j++ {
	// 	fmt.Println(j + 1)

	// }

	var input string
	for {
		fmt.Print("Enter a string : ")
		fmt.Scanf("%s \n", &input)
		if input == "ex" {
			break
		}
	}

}
