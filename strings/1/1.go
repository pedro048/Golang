package main

import(
	"fmt"
	"strings"
)
func main() {
	text := "A text file is a file with a lot of text in it."
	findText := "text"
	modText := "doc"
	aplitText := " "
	
	fmt.Println(strings.Contains(text, findText))
	fmt.Println(strings.Count(text, findText))
	fmt.Println(strings.Index(text, findText))
	fmt.Println(strings.ReplaceAll(text, findText, modText))
	fmt.Println(strings.Split(text, aplitTrimText))
	fmt.Println(strings.ToLower(text))
	fmt.Println(strings.ToUpper(text))
	fmt.Println(strings.Trim(text, aplitTrimText))
			
}