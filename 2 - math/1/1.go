package main

import (
	"fmt"
)

func main() {
	var a = 3.1412

	res_1 := math.ceil(a)
	fmt.Println(res_1)

	res_2 := math.floor(a)
	fmt.Println(res_2)

	res_3 := math.round(a)
	fmt.Println(res_3)

	res_4 := math.trunc(a)
	fmt.Println(res_4)

}
