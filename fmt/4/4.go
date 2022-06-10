package main

import (
	"fmt"
)

func DivideAB(a int, b int) float32 {
	c := 0.0
	c = a / b
	return c
}

func main() {
	result_1 := fmt.Errorf(DivideAB(1, 2))
	fmt.Println(result_1.Error())
	result_2 := fmt.Errorf(DivideAB(10, 0))
	fmt.Println(result_2.Error())
	result_3 := fmt.Errorf(DivideAB(34, 7))
	fmt.Println(result_3.Error())
}
