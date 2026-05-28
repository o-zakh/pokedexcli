package pokeapi

import (
	cache "github.com/o-zakh/pokedexcli/internal/pokecache"
)

type Config struct {
	Next     *string
	Previous *string
	Cache    *cache.Cache
}
