package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/xfac11/pokedexcli/internal/poketypes"
)

func commandCatch(c *config, args []string) error {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", args[0])
	var pokemon poketypes.Pokemon
	data, ok := c.Cache.Get(url)
	if ok {
		if err := json.Unmarshal(data, &pokemon); err != nil {
			return err
		}
	} else {
		request, err := http.NewRequest("GET", url, nil)
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
			return fmt.Errorf("Could not create decoder")
		}

		decoder.Decode(&pokemon)
		if data, err := json.Marshal(&pokemon); err == nil {
			c.Cache.Add(url, data)
		} else {
			return err
		}
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])
	if !tryCatch(pokemon.BaseExperience) {
		fmt.Printf("%s escaped!\n", args[0])
		return nil
	}

	c.Pokedex[args[0]] = pokemon
	fmt.Printf("%s was caught!\n", args[0])
	return nil
}

// < 1xx = 100% 1xx = 50% 2xx 33% 3xx 25% 4xx 20% and so on
func tryCatch(baseExperience int) bool {
	n := int(float32(baseExperience) * 0.01)
	if n == 0 {
		return true
	}
	return rand.Intn(n+1) == 0
}
