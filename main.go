package main

import (
	"bufio"
	"encoding/json"
	"fmt"
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
		"inspect": {
			name:        "inspect",
			description: "Displays information about the given pokemon in your pokedex",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Display all your caught pokemon",
			callback:    commandPokedex,
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
