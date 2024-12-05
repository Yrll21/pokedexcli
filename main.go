package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name string
	description string
	callback func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
	}
}

func main() {
	for {
		fmt.Print("pokedex > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		switch input {
		case "help":
			commandHelp()
		case "exit":
			commandExit()
		default:
			fmt.Println("Unknown command:", input)
		}
		}
}

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}