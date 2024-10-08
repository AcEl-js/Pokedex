package main

import "fmt"

func callBackMypokemons(cfg *config,mapArea string, name string) error {
    fmt.Println("your Pokemons:")
    for _,pokemon := range cfg.myPokemons{
        fmt.Println("- ",pokemon)
    }
return nil
}
