package main

import (
	"fmt"
	"os"

	pk "github.com/o-zakh/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(config *pk.Config, param string) error
}

func commandsMap() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world. Each subsequent call to map displays the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "It's similar to the 'map' command, however, instead of displaying the next 20 locations, it displays the previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Shows a list of all the Pokémon located in the location specified after the command. ex: explore pastoria-city-area",
			callback:    commandExplore,
		},
	}
}

func commandExit(config *pk.Config, param string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *pk.Config, param string) error {
	fmt.Printf("\nWelcome to the Pokedex!\nUsage:\n")
	for _, value := range commandsMap() {
		fmt.Printf("\n%v: %v", value.name, value.description)
	}
	fmt.Println()
	return nil
}

func commandMap(config *pk.Config, param string) error {
	pk.Pokeapi_LocAreaForward(config)
	return nil
}

func commandMapb(config *pk.Config, param string) error {
	pk.Pokeapi_LocAreaBack(config)
	return nil
}

func commandExplore(config *pk.Config, param string) error {
	pk.Pokeapi_ExpNameList(config, param)
	return nil
}
