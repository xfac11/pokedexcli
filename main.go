package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type config struct {
	Next      string
	Previous  string
	FirstTime bool
}
type cliCommand struct {
	name        string
	description string
	callback    func(config *config) error
}
type locationArea struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
type locationAreaResponse struct {
	Count    int            `json:"Count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Resulst  []locationArea `json:"results"`
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
	}
	return commandMap
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	config := config{
		FirstTime: true,
	}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)
		command, ok := getCommands()[cleaned[0]]
		if !ok {
			fmt.Printf("Unknown command: %s\n", cleaned)
			continue
		}
		err := command.callback(&config)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func commandHelp(c *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	for key, value := range getCommands() {
		fmt.Printf("%s: %s\n", key, value.description)
	}
	return nil
}

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(c *config) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if c.Next != "" {
		url = c.Next
	}
	if !c.FirstTime && c.Next == "" {
		println("you're on the last page")
		return nil
	}

	c.FirstTime = false

	var location_areas locationAreaResponse
	err := getLocationAreas(url, &location_areas)
	if err != nil {
		return nil
	}

	for i := range 20 {
		println(location_areas.Resulst[i].Name)
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

func commandMapb(c *config) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if c.Previous != "" {
		url = c.Previous
	}
	if c.Previous == "" {
		println("you're on the first page")
		return nil
	}
	c.FirstTime = false
	var location_areas locationAreaResponse
	err := getLocationAreas(url, &location_areas)
	if err != nil {
		return err
	}

	for i := range 20 {
		println(location_areas.Resulst[i].Name)
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

func getLocationAreas(url string, locationAreas *locationAreaResponse) error {
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
