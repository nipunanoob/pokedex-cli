package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"net/http"
	"io"
	"encoding/json"
)

type cliCommand struct {
	name	string
	description	string
	callback	func() error
	LocationResponse	*LocationResponse
}

type LocationResponse struct {
	Results	[]Location	`json:"results"`
	Next *string	`json:"next"`
	Previous	*string	`json:"previous"`
}

type Location struct {
	Name	string	`json:"name"`
}

var commandList map[string]cliCommand

func initializeCommands() map[string]cliCommand{
	nextURL := "https://pokeapi.co/api/v2/location-area/"
	previousURL := ""
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
		"map": {
			name: "map",
			description: "Displays 20 location areas in Pokemon. Each call will display next 20 locations",
			callback:  commandMap,
			LocationResponse: &LocationResponse{
				Next: &nextURL,
				Previous: &previousURL,	
			},
		},
	}
	return commands
}

func main() {

	commandList = initializeCommands()
	scanner := bufio.NewScanner(os.Stdin) // start scanner to read user standard input
	for {
		fmt.Print("Pokedex > ") 
		scanner.Scan() //scans for next line
		command := cleanInput(scanner.Text())[0]
		if command == ""{
			continue
		}
		val, ok := commandList[command]
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
	for _,v := range commandList{
		fmt.Printf("%s: %s\n",v.name,v.description)
	}
	return nil
}

func commandMap() error {
	mapCmd := commandList["map"]

	if mapCmd.LocationResponse == nil || mapCmd.LocationResponse.Next == nil {
		return fmt.Errorf("No next URL found")
	}
	res, err := http.Get(*mapCmd.LocationResponse.Next)
	if err != nil {
		return fmt.Errorf("Error: %w", err)
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("Response not successful: %d", res.StatusCode)
	}
	if err != nil {
		return fmt.Errorf("Error: %w", err)
	}

	var locRes LocationResponse
	err = json.Unmarshal(body, &locRes)
	if err != nil {
		return fmt.Errorf("Error unmarshalling: %w", err)
	}
	for _, loc := range locRes.Results {
        fmt.Println(loc.Name)
    }
	
	if locRes.Next != nil {
		mapCmd.LocationResponse.Next = locRes.Next
	} else {
		mapCmd.LocationResponse.Next = nil
	}

	if locRes.Previous != nil {
		mapCmd.LocationResponse.Previous = locRes.Previous
	} else {
		mapCmd.LocationResponse.Previous = nil
	}
	
	return nil
}