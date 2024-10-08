package main

import (
	"time"

	pokeapi "github.com/AcEl-js/Pokedex/internal/pokesipa"
)


type config struct{
    pokeapiClient pokeapi.Client
    nextLocationUrl *string
    prevousLocationUrl *string
    myPokemons []string

}

func main () {
    cfg := config{
        pokeapiClient: pokeapi.NewClient(5*time.Minute),
        myPokemons: []string{},
    }
     startRepl(&cfg)
}


