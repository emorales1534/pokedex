package main

import(
	"fmt"
	"errors"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("provide one location")
	}
	location := args[0]

	resp, err := cfg.pokeapiClient.ExploreLocation(location)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s\n", location)
	fmt.Printf("Found Pokemon:\n")
	for _, encounter := range resp.PokemonEncounters {
		fmt.Println("- %s\n", encounter.Pokemon.Name)
	}

	return nil
}
