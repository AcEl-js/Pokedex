package pokeapi

import (
	"encoding/json"
	"net/http"
)
func (c *Client) GetPokemonInfo(name string) (int,error){

    endpoint := "/pokemon/"
    fullPath := baseUrl + endpoint +name


    req, err := http.NewRequest("GET",fullPath,nil)
    if err != nil {
        return 0, err
    }

    res, err := c.httpClient.Do(req)
    if err != nil {
        return 0,err 
    }

    defer res.Body.Close()

    var pokemonInfo Pokemoninf
   if  err := json.NewDecoder(res.Body).Decode(&pokemonInfo);err != nil{
        return 0,nil
    }

    return pokemonInfo.BaseExperience,nil
}
