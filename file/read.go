package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	count := 1

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Line ", count, ":", line)
		count++
	}

}
