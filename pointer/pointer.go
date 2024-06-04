package main

import "fmt"

func zeroval(i int) {
	i = 0
}

func zeroptr(i *int) {
	*i = 0
}

func main() {
	i := 1
	fmt.Println("Initial value of i:", i)
	zeroval(i)
	fmt.Println("Value of i after zeroval(i):", i)
	zeroptr(&i)
	fmt.Println("Value of i after zeroptr(&i):", i)
	fmt.Println("Pointer of i:", &i)
}
