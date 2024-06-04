package main

import "fmt"

func main() {
	var input string
	var output string

	for {
		fmt.Print("Enter a color : ")
		fmt.Scanf("%s \n", &input)
		switch input {
		case "blue":
			output = "#0000FF"
		case "green":
			output = "#00FF00"
		case "pink":
			output = "#FF00FF"
		default:
			output = "#FFFF00"
		}
		fmt.Println("Hex code for", input, "is", output)
		if input == "ex" {
			break
		}
	}

}
