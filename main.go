package main

import (
	"time"

	"github.com/emorales1534/pokedex/internal/pokeapi"
)


func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second, 5 * time.Minute) 
	cfg := &config {
		commands: getCommands(),
		pokeapiClient: pokeClient,
		pokedex: make(map[string]pokeapi.RespPokemon),
	}

	startRepl(cfg)
}
