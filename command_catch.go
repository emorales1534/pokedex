package main

import(
  "fmt"
  "errors"
  "math/rand"
  )

func commandCatch(cfg *config, args ...string) error {
  if len(args) != 1 {
    return errors.New("provide one pokemon")
  }

  pokemon := args[0]
 
  fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)

  resp, err := cfg.pokeapiClient.CatchPokemon(pokemon)
  if err != nil {
    return err
  }

  baseExp := resp.BaseExperience
  
  chance := 0

  if baseExp >= 1 {
    chance = rand.Intn(baseExp)
  }

  caught := chance < 30

  if !caught {
    fmt.Printf("%s escaped!\n", pokemon)
  } else {
    fmt.Printf("%s was caught!\n", pokemon)
    cfg.pokedex[pokemon] = resp
  }
 
  return nil

}
