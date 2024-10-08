package main

import "fmt"


func callBackExplore(cfg *config,mapArea string) error{
    pokemons,err:= cfg.pokeapiClient.ExploreLocation(mapArea)
    if err != nil{
        return err
    }
    for _, pokemoneName := range pokemons {

        fmt.Println(pokemoneName)
        

    }
    return nil
}
