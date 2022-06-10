package main

import "fmt"

func main() {
	var a int = 3
	const b = 3.1412
	format := "The first number is %d and the second number is %.2f."
	myOutput := fmt.Sprintf(format, a, b)
	fmt.Printf(myOutput)

}
