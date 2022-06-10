package main

import (
	"fmt"
)

func main() {
	var b = 3.5
	var c = 4.5
	var d = 8.78

	res_1 := math.ceil(b)
	fmt.Println(res_1)

	res_2 := math.floor(b)
	fmt.Println(res_2)

	res_3 := math.round(b)
	fmt.Println(res_3)

	res_4 := math.trunc(b)
	fmt.Println(res_4)

	res_5 := math.ceil(c)
	fmt.Println(res_5)

	res_6 := math.floor(c)
	fmt.Println(res_6)

	res_7 := math.round(c)
	fmt.Println(res_7)

	res_8 := math.trunc(c)
	fmt.Println(res_8)

	res_9 := math.ceil(d)
	fmt.Println(res_9)

	res_10 := math.floor(d)
	fmt.Println(res_10)

	res_11 := math.round(d)
	fmt.Println(res_11)

	res_12 := math.trunc(d)
	fmt.Println(res_12)

}
