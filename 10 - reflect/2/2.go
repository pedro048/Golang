package main

import (
	"fmt"
	"reflect"
)

func MultipleArguments(a int, b string, c bool, d int64) {
	var aux_1 int = 1
	var aux_2 int64 = int64(2222)
	if reflect.TypeOf(a) == reflect.TypeOf(aux_1) || reflect.TypeOf(a) == reflect.TypeOf(aux_2) {
		a = a + 1
	}
	if reflect.TypeOf(b) == reflect.TypeOf(aux_1) || reflect.TypeOf(b) == reflect.TypeOf(aux_2) {

	}
	if reflect.TypeOf(c) == reflect.TypeOf(aux_1) || reflect.TypeOf(c) == reflect.TypeOf(aux_2) {

	}
	if reflect.TypeOf(d) == reflect.TypeOf(aux_1) || reflect.TypeOf(d) == reflect.TypeOf(aux_2) {
		d = d + 1
	}
	fmt.Printf("The value of a is %d, the value of b is %s, the value of c is %v and the value of d is %d. ", a, b, c, d)
}

func main() {
	var a int = 32
	b := "Text"
	var c bool = true
	var d int64 = int64(1235)
	fmt.Printf("The value of a is %d, the value of b is %s, the value of c is %v and the value of d is %d. ", a, b, c, d)
	fmt.Printf("\n")
	MultipleArguments(a, b, c, d)

}
