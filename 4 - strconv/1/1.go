package main

import (
	"fmt"
	"strconv"
)

func main() {
	var boolValue bool = true
	var intValue int = 48
	var floatValue float64 = 3.1415

	s_1 := strconv.Itoa(intValue)
	s_2 := strconv.FormatBool(boolValue)
	s_3 := strconv.FormatFloat(floatValue, 'E', -1, 64)
	text := "The boolText is " + s_2 + " the intText is " + s_1 + " and the floatText is " + s_3

	fmt.Printf(text)

}
