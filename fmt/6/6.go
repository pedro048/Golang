package main

import "fmt"

var x int = 32

func PlusOne(x int) int {
	x = x + 1
	return x
}

func main() {
	fmt.Println(x)
	PlusOne(x)
	fmt.Println(x)
}
