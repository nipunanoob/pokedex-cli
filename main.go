package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

type cliCommand struct {
	name	string
	description	string
	callback	func() error
}

var commandMap map[string]cliCommand

func initializeCommands() map[string]cliCommand{
	commands := map[string]cliCommand {
		"exit": {
			name: "exit",
			description: "Exit the pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback:  commandHelp,
		},
	}
	return commands
}

func main() {

	commandMap = initializeCommands()
	scanner := bufio.NewScanner(os.Stdin) // start scanner to read user standard input
	for {
		fmt.Print("Pokedex > ") 
		scanner.Scan() //scans for next line
		command := cleanInput(scanner.Text())[0]
		if command == ""{
			continue
		}
		val, ok := commandMap[command]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			val.callback()
		}
	}
}

func cleanInput(text string) []string {

	// gets slice of words in text after trimming, lowercasing and seperating string by whitespace

	if text == ""{
		return []string{""}
	}
	formattedText := strings.Split(strings.ToLower(strings.Trim(text, " ")), " ")
	return formattedText
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!\nUsage\n")
	for _,v := range commandMap{
		fmt.Printf("%s: %s\n",v.name,v.description)
	}
	return nil
}