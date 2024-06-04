package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getInput(promt string) float64 {
	fmt.Printf("%v", promt) // %v is a placeholder for the prompt string in
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		message, _ := fmt.Scanf("%v must be a number", promt)
		panic(message)
	}
	return value
}

func getOpereator() string {
	fmt.Print("Enter an operator (+, -, *, /) : ")
	op, _ := reader.ReadString('\n')
	return strings.TrimSpace(op)
}

func add(value1, value2 float64) float64 {
	return value1 + value2
}

func subtract(value1, value2 float64) float64 {
	return value1 - value2
}

func multiply(value1, value2 float64) float64 {
	return value1 * value2
}

func divide(value1, value2 float64) float64 {
	return value1 / value2
}

func main() {
	value1 := getInput("Enter first number : ")
	value2 := getInput("Enter second number : ")

	var result float64

	switch operator := getOpereator(); operator {
	case "+":
		result = add(value1, value2)
	case "-":
		result = subtract(value1, value2)
	case "*":
		result = multiply(value1, value2)
	case "/":
		result = divide(value1, value2)
	default:
		panic("Invalid operator")
	}
	fmt.Printf("Result is %v", result)
}
