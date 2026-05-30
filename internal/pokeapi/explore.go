package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type ExpLoc struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

// {
// 	PokemonEncounters []struct {
// 		Pokemon struct {
// 			Name string `json:"name"`
// 			URL  string `json:"url"`
// 		} `json:"pokemon"`
// 	} `json:"pokemon_encounters"`
// }

func Pokeapi_ExpResponse(config *Config, param string) ExpLoc {
	if param == "" { // Надо выяснить, какое значение у param, если cleaned[1] просто нет, в консоль ввели только 1 слово (команду)
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
