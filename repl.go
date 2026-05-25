package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	loweredText := strings.ToLower(text)
	output := strings.Fields(loweredText)
	return output
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		var input string
		if scanner.Scan() {
			input = scanner.Text()
		}
		cleaned := cleanInput(input)
		command, ok := commandMap()[cleaned[0]]
		if ok {
			err := command.callback()
			if err != nil {
				fmt.Println(command.callback())
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}
