package main

import "fmt"

func commandInspect(c *config, args []string) error {
	if args == nil {
		return fmt.Errorf("Command inspect requires a pokemon as argument. Got nil")
	}
	if len(args) == 0 || len(args) > 2 {
		return fmt.Errorf("Command inspect requires one pokemon as argument. Got zero or more")
	}
	data, ok := c.Pokedex[args[0]]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", data.Name)
	fmt.Printf("Height: %d dm\n", data.Height)
	fmt.Printf("Weight: %d hg\n", data.Weight)
	fmt.Printf("Stats: \n")
	for _, stat := range data.Stats {
		fmt.Printf(" -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, typ := range data.Types {
		fmt.Printf(" - %s\n", typ.Type.Name)
	}
	return nil
}
