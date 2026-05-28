package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	pk "github.com/o-zakh/pokedexcli/internal/pokeapi"
)

func cleanInput(text string) []string {
	loweredText := strings.ToLower(text)
	output := strings.Fields(loweredText)
	return output
}

func startRepl(config *pk.Config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		var input string
		if scanner.Scan() {
			input = scanner.Text()
		}
		cleaned := cleanInput(input)
		command, ok := commandsMap(config)[cleaned[0]]
		if ok {
			err := command.callback(config)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}
