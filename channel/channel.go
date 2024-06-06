package main

import (
	"fmt"
	"time"
)

func process1(c chan string, data string) {
	c <- data
}

func main() {
	ch := make(chan string)
	go process1(ch, "Hello World")
	fmt.Println("Data from channel: ", <-ch)
	time.Sleep(5 * time.Second)

}
