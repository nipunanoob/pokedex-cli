package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println("Hello, World!")
	fmt.Println(cleanInput(" hello world test 123         "))
}

func cleanInput(text string) []string {
	if text == ""{
		return nil
	}
	formattedText := strings.Split(strings.Trim(text, " "), " ")
	return formattedText
}