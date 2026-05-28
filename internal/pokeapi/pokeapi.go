package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type LocArea struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func PokeApi_LocAreaResponse(config *Config, link *string) LocArea {
	if link == nil {
		fmt.Println("\nYou reached the end of the list\n")
		return LocArea{}
	}
	res, err := http.Get(*link)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	var locArea LocArea
	err = json.Unmarshal(body, &locArea)
	if err != nil {
		log.Fatal(err)
	}
	config.Next = locArea.Next
	config.Previous = locArea.Previous
	return locArea
}

func Pokeapi_LocAreaForward(config *Config) {
	locArea := PokeApi_LocAreaResponse(config, config.Next)
	Pokeapi_LocAreaNameList(locArea)
}

func Pokeapi_LocAreaBack(config *Config) {
	locArea := PokeApi_LocAreaResponse(config, config.Previous)
	Pokeapi_LocAreaNameList(locArea)
}

func Pokeapi_LocAreaNameList(locArea LocArea) {
	for _, location := range locArea.Results {
		fmt.Println(location.Name)
	}
}
