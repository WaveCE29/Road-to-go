package main

import (
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data1 := []byte("Hello, World!")
	err := os.WriteFile("data.txt", data1, 0644)
	if err != nil {
		panic(err)
	}

	f, err := os.Create("emplyeeName")
	check(err)

	defer f.Close()

	data2 := []byte("John Doe")
	os.WriteFile("emplyeeName.txt", data2, 0644)
}
