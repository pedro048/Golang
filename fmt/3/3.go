package main

import (
	"fmt"
)

func main() {

	// Declaring two variables
	var a int = 3
	const b = 3.1412

	format := "The first number is %d and the second number is %f."
	myOutput := fmt.Sprintf(format, a, b)
	fmt.Printf(myOutput)

	fmt.Printf("\n")

	var first int
	var second float32

	fmt.Sscanf("3;3.1412", "%d;%f", &first, &second)
	fmt.Printf(format, first, second)
}
