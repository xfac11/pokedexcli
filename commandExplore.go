package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/xfac11/pokedexcli/internal/poketypes"
)

func commandExplore(c *config, args []string) error {
	if args == nil {
		return fmt.Errorf("Command explore requires a location area as argument. Got nil")
	}
	if len(args) == 0 || len(args) > 2 {
		return fmt.Errorf("Command explore requires one location area as argument. Got zero or more")
	}
	fullUrl := "https://pokeapi.co/api/v2/location-area/" + args[0] + "/"
	data, ok := c.Cache.Get(fullUrl)
	var area poketypes.LocationArea
	if !ok {
		request, err := http.NewRequest("GET", fullUrl, nil)
		if err != nil {
			return err
		}
		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			return err
		}
		defer response.Body.Close()

		if response.StatusCode != 200 {
			return fmt.Errorf("No pokemon found with that name, status : %s", response.Status)
		}

		decoder := json.NewDecoder(response.Body)
		if decoder == nil {
			return fmt.Errorf("Decoder was nil")
		}

		err = decoder.Decode(&area)
		if err != nil {
			return err
		}

		jsonData, err := json.Marshal(area)
		if err != nil {
			return err
		}
		err = c.Cache.Add(fullUrl, jsonData)
		if err != nil {
			return err
		}
	} else {
		json.Unmarshal(data, &area)
	}

	for _, encounter := range area.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}
