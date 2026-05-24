package main

import (
	"strings"
)

func cleanInput(text string) []string {
	loweredText := strings.ToLower(text)
	output := strings.Fields(loweredText)

	// fmt.Printf("loweredText: %v\n", loweredText)
	// fmt.Printf("Output: %v\n", output)
	// fmt.Printf("Output[0]: %v\n", output[0])
	return output
}
