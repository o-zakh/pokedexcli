package main

import pk "github.com/o-zakh/pokedexcli/internal/pokeapi"

func main() {
	url := "https://pokeapi.co/api/v2/location-area/"
	config := pk.Config{
		Next: &url,
	}
	startRepl(&config)
}
