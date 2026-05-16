package main

import (
	"encoding/json"
	"fmt"

	"github.com/xfac11/pokedexcli/internal/poketypes"
)

func commandMap(c *config, args []string) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if c.Next != "" {
		url = c.Next
	}
	if !c.FirstTime && c.Next == "" {
		fmt.Println("you're on the last page")
		return nil
	}

	c.FirstTime = false

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
		println(location_areas.Results[i].Name)
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
