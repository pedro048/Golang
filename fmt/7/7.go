package main

import "fmt"

var x int = 32
var v *int

func PlusOne(v *int) {
	v = &x
	*v = *v + 1
}

func main() {
	fmt.Println(x)
	PlusOne(v)
	fmt.Println(x)
}
