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
			description: "Exit the Pokédex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokémon world. Each subsequent call to map displays the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "It's similar to the 'map' command, however, instead of displaying the next 20 locations, it displays the previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Shows a list of all the Pokémon located in the location specified as an argument. ex: explore pastoria-city-area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Make an attempt to catch a Pokémon specified as an argument. ex: catch pikachu",
			callback:    commandCatch,
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

func commandCatch(config *pk.Config, param string) error {
	if param == "" {
		fmt.Println("Input the name of the Pokémon you would like to catch as an argument. ex: catch pikachu")
		return nil
	}
	pokemon, caught := pk.Pokeapi_CatchAttempt(config, param)
	if caught {
		config.Pokedex[param] = pokemon
		fmt.Printf("%v was caught!", param)
		fmt.Println()
	} else {
		fmt.Printf("%v escaped!", param)
		fmt.Println()
	}
	return nil
}
