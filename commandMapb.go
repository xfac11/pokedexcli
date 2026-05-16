package main

import (
	"encoding/json"
	"fmt"

	"github.com/xfac11/pokedexcli/internal/poketypes"
)

func commandMapb(c *config, args []string) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if c.Previous != "" {
		url = c.Previous
	}
	if c.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	var location_areas poketypes.LocationAreaSet
	data, ok := c.Cache.Get(url)
	if ok {
		json.Unmarshal(data, &location_areas)
	} else {
		err := requestLocationAreas(url, &location_areas)
		jsonData, err := json.Marshal(location_areas)
		if err != nil {
			return err
		}
		err = c.Cache.Add(url, jsonData)
		if err != nil {
			return err
		}
	}

	for i := range 20 {
		fmt.Println(location_areas.Results[i].Name)
	}

	if location_areas.Next != nil {
		c.Next = *location_areas.Next
	} else {
		c.Next = ""
	}
	if location_areas.Previous != nil {
		c.Previous = *location_areas.Previous
	} else {
		c.Previous = ""
	}

	return nil
}
