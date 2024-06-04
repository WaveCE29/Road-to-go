package main

import "fmt"

var grade int

func main() {
	fmt.Print("Enter your grade:")
	fmt.Scanf("%d", &grade)
	if grade >= 80 {
		fmt.Println("Grade A")
	} else if grade >= 70 {
		fmt.Println("Grade B")
	} else if grade >= 60 {
		fmt.Println("Grade C")
	} else if grade >= 50 {
		fmt.Println("Grade D")
	} else {
		fmt.Println("Grade F")
	}

}
