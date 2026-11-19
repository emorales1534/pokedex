package main

import(
	"fmt"
	"errors"
)

//Inspecting pokemon: only caught and indexed
func commandInspect (cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("provide one pokemon")
	}

	name := args[0]

	pokemon, ok := cfg.pokedex[name]

	if ok == false {
		return errors.New("you have not caught that pokemon")
	} else {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %v\n", pokemon.Height)
		fmt.Printf("Weight: %v\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, typeInfo := range pokemon.Types {
			fmt.Println("  -", typeInfo.Type.Name)
		}
	}

	return nil
}
