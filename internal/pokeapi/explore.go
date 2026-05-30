package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type ExpLoc struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func Pokeapi_ExpResponse(config *Config, param string) ExpLoc {
	if param == "" {
		fmt.Println("Input the explored location name as an argument. ex: explore pastoria-city-area")
		return ExpLoc{}
	}

	link := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%v", param)
	body, exists := config.Cache.Get(link)

	if !exists {
		res, err := http.Get(link)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		body, err = io.ReadAll(res.Body)
		if res.StatusCode == 404 {
			fmt.Println("Can't find the location in the database. Try looking for a correct location name using the command 'map'")
			return ExpLoc{}
		}
		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}
		if err != nil {
			log.Fatal(err)
		}
	}

	config.Cache.Add(link, body)

	var expLoc ExpLoc
	err := json.Unmarshal(body, &expLoc)
	if err != nil {
		log.Fatal(err)
	}
	return expLoc
}

func Pokeapi_ExpNameList(config *Config, param string) {
	for _, encounter := range Pokeapi_ExpResponse(config, param).PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
}
