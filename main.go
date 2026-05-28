package main

import (
	"time"

	pk "github.com/o-zakh/pokedexcli/internal/pokeapi"
	cache "github.com/o-zakh/pokedexcli/internal/pokecache"
)

func main() {

	url := "https://pokeapi.co/api/v2/location-area/"
	config := pk.Config{
		Next:  &url,
		Cache: cache.NewCache(5 * time.Second),
	}
	startRepl(&config)
}
