package main

import (
	"fmt"
	"strconv"
)

func main() {
	var boolText bool = true
	var intText int = 48
	var floatText float64 = 3.1415

	intValue := strconv.Itoa(intText)
	boolValue := strconv.FormatBool(boolText)
	floatValue := strconv.FormatFloat(floatText, 'E', -1, 64)
	format := "The boolValue is %v, the intValue is %s and the floatValue is %s"

	fmt.Printf(format, boolValue, intValue, floatValue)

}
