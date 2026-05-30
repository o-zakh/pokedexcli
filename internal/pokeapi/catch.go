package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
)

type PokemonInfo struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

func Pokeapi_CatchResponse(config *Config, param string) PokemonInfo {

	link := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%v", param)
	body, exists := config.Cache.Get(link)

	if !exists {
		res, err := http.Get(link)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		body, err = io.ReadAll(res.Body)
		if res.StatusCode == 404 {
			fmt.Println("Can't find the Pokémon in the database. Try searching for an existing Pokémon in the wilds using the 'explore' command")
			return PokemonInfo{}
		}
		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}
		if err != nil {
			log.Fatal(err)
		}
	}

	config.Cache.Add(link, body)

	var pokemonInfo PokemonInfo
	err := json.Unmarshal(body, &pokemonInfo)
	if err != nil {
		log.Fatal(err)
	}
	return pokemonInfo
}

func Pokeapi_CatchAttempt(config *Config, param string) (PokemonInfo, bool) {
	// Formula: baseExp - ((baseExp - 50) / 2)
	// Example: baseExp = 166. Any random number in range 0-108 wins, 109 - 166 — loses.

	fmt.Printf("Throwing a Pokeball at %v...", param)
	fmt.Println()
	autowin := 50
	pokemon := Pokeapi_CatchResponse(config, param)
	baseExp := pokemon.BaseExperience
	baseExp -= autowin
	if baseExp <= 0 {
		return pokemon, true
	}

	randomNum := rand.Intn(baseExp)
	if randomNum >= baseExp-((baseExp-autowin)/2) {
		return PokemonInfo{}, false
	} else {
		return pokemon, true
	}
}
