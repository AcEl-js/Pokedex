package pokeapi

import (
	"encoding/json"
	"net/http"
)


func (c *Client) ExploreLocation(mapArea string) ([]string,error){

    endpoint := "/location-area/"
    fullUrl :=baseUrl+ endpoint + mapArea
    req, err := http.NewRequest("GET",fullUrl,nil)
    if err != nil {
        return []string{},err
    }

    res, err := c.httpClient.Do(req)

    if err != nil {
        return []string{},err
    }
    defer res.Body.Close()

    var pokemons Pokemons
    var pokemonsList []string

    if err := json.NewDecoder(res.Body).Decode(&pokemons); err!= nil {
        return []string{},err
    }
    for _,k := range pokemons.PokemonEncounters{

        pokemonsList = append(pokemonsList, k.Pokemon.Name)
    }
    return pokemonsList, nil
}
