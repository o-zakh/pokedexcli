package main

import (
	"bufio"
	"fmt"
	"log"
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
		fmt.Print("\nPokedex > ")
		var input string
		if scanner.Scan() {
			input = scanner.Text()
			if err := scanner.Err(); err != nil {
				log.Fatalf("reading stdin: %v", err)
			}
		}
		cleaned := cleanInput(input)
		if len(cleaned) == 0 {
			fmt.Println("Input a command. Enter 'help' to an available commands list")
			continue
		}
		command, ok := commandsMap()[cleaned[0]]
		if ok {
			param := ""
			if len(cleaned) > 1 {
				param = cleaned[1]
			}
			err := command.callback(config, param)
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
