package main

import (
	"fmt"
)

func main() {
	var a int = 3
	const b = 3.1412
	format := "The first number is %d and the second number is %.2f."
	fmt.Printf(format, a, b)
}
