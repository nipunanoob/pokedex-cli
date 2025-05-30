package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin) // start scanner to read user standard input
	for {
		fmt.Print("Pokedex > ") 
		scanner.Scan() //scans for next line
		text := cleanInput(scanner.Text())
		fmt.Printf("Your command was: %s\n", text[0])
	}
}

func cleanInput(text string) []string {

	// gets slice of words in text after trimming, lowercasing and seperating string by whitespace

	if text == ""{
		return nil
	}
	formattedText := strings.Split(strings.ToLower(strings.Trim(text, " ")), " ")
	return formattedText
}