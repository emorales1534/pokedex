package pokeapi

import (
  "encoding/json"
  "io"
  "net/http"
  )

//catch pokemon
func (c *Client) CatchPokemon(pokemon string) (RespPokemon, error) {
  url := baseURL + "/pokemon/" + pokemon

  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
	  return RespPokemon{}, err
  }

  resp, err := c.httpClient.Do(req)
  if err != nil {
	  return RespPokemon{}, err
  }
  defer resp.Body.Close()

  dat, err := io.ReadAll(resp.Body)
  if err != nil {
	  return RespPokemon{}, err
  }

  pokemonResp := RespPokemon{}
  err = json.Unmarshal(dat, &pokemonResp)
  if err != nil {
	  return RespPokemon{}, err
  }

  return pokemonResp, nil
  
}
