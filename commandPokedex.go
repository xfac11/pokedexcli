package main

import "fmt"

func commandPokedex(c *config, args []string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range c.Pokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}
