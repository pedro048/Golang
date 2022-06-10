package main

import (
	"fmt"
	"reflect"
)

func MultipleArguments(x interface{}, y interface{}, z interface{}, w interface{}) {
	fmt.Println("The type of arg 0 is", reflect.TypeOf(x))
	fmt.Println("The type of arg 1 is", reflect.TypeOf(y))
	fmt.Println("The type of arg 2 is", reflect.TypeOf(z))
	fmt.Println("The type of arg 3 is", reflect.TypeOf(w))
}

func main() {
	MultipleArguments(34, "This is a string", 3.14, " ")
	MultipleArguments(int64(34), "This is a string", float32(3.14), " ")
	MultipleArguments(int64(10000000000000), "This is a string", 345.45645634634, true)
}