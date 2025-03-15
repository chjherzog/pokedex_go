package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CleanInput(text string) []string {
	trimmedText := strings.Fields(text)
	var loweredText []string
	for i := range trimmedText {
		loweredText = append(loweredText, strings.ToLower(trimmedText[i]))
	}
	return loweredText
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(allCliCommands map[string]cliCommand) error {
	fmt.Println("Available commands")
	for _, command := range allCliCommands {
		fmt.Sprintf("%s: %s", command.name, command.description)
	}
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	allCliCommands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exits the program",
			callback:    commandExit,
		},
	}
	allCliCommands["help"] = cliCommand{
		name:        "help",
		description: "Exits the program",
		callback:    func() error { return commandHelp(allCliCommands) },
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	for {
		fmt.Println()
		for _, val := range allCliCommands {
			fmt.Println(val.name + ": " + val.description)
		}
		scanner.Scan()
		line := CleanInput(scanner.Text())

		if len(line) == 0 {
			continue
		}

		firstWord := line[0]

		command, exists := allCliCommands[firstWord]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		if err := command.callback(); err != nil {
			fmt.Println("Error: ", err)
		}
	}
}
