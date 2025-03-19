package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	config := &Config{}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		line := CleanInput(scanner.Text())
		if len(line) == 0 {
			continue
		}

		commandName := line[0]
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(config)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown Command")
			continue
		}

	}
}

func CleanInput(text string) []string {
	trimmedText := strings.Fields(text)
	var loweredText []string
	for i := range trimmedText {
		loweredText = append(loweredText, strings.ToLower(trimmedText[i]))
	}
	return loweredText
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exits the program",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Exits the program",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Lists twenty locations",
			callback:    commandGet,
		},
	}
}
