package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/xfac11/pokedexcli/internal/pokecache"
	"github.com/xfac11/pokedexcli/internal/poketypes"
	"github.com/xfac11/pokedexcli/internal/repl"
)

type config struct {
	Next      string
	Previous  string
	FirstTime bool
	Cache     *pokecache.Cache
	Pokedex   map[string]poketypes.Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(config *config, args []string) error
}

func getCommands() map[string]cliCommand {
	commandMap := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays 20 location areas. subsequent call to this command displays the next 20 areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "bmap",
			description: "Displays the previous 20 location areas. subsequent call to this command displays the previous 20 areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Displays the given areas possible encounters",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Throws a pokeball and tries to catch the pokemon",
			callback:    commandCatch,
		},
	}
	return commandMap
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	config := config{
		FirstTime: true,
		Cache:     pokecache.NewCache(time.Second * 5),
		Pokedex:   make(map[string]poketypes.Pokemon),
	}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleaned := repl.CleanInput(text)
		command, ok := getCommands()[cleaned[0]]
		if !ok {
			fmt.Printf("Unknown command: %s\n", cleaned)
			continue
		}
		err := command.callback(&config, cleaned[1:])
		if err != nil {
			fmt.Println(err)
		}
	}
}

func commandHelp(c *config, args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	for key, value := range getCommands() {
		fmt.Printf("%s: %s\n", key, value.description)
	}
	return nil
}

func commandExit(c *config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

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

	fmt.Printf("Throwing a Pokeball at %s...", args[0])
	if !tryCatch(pokemon.BaseExperience) {
		fmt.Printf("%s escaped!\n", args[0])
		return nil
	}

	c.Pokedex[args[0]] = pokemon
	fmt.Printf("%s was caught!\n", args[0])
	return nil
}

func tryCatch(baseExperience int) bool {
	return rand.Intn(int(float32(baseExperience)*0.1)) == 1
}

func requestLocationAreas(url string, locationAreas *poketypes.LocationAreaSet) error {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	if response.StatusCode != 200 {
		return fmt.Errorf("Non 200 status code: %s", response.Status)
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	if decoder == nil {
		return fmt.Errorf("Could not create decoder")
	}
	err = decoder.Decode(locationAreas)
	if err != nil {
		return err
	}
	return nil
}
