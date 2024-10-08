package main

import pokeapi "github.com/AcEl-js/Pokedex/pokesipa"


type config struct{
    pokeapiClient pokeapi.Client
    nextLocationUrl *string
    prevousLocationUrl *string

}

func main () {
    cfg := config{
        pokeapiClient: pokeapi.NewClient(),

    }
     startRepl(&cfg)
    


}


