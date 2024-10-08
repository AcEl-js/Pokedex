package main

import (
	"fmt"
	"math/rand"
	"time"

)


func randomBool(dificulty float64) bool{
   rand.Seed(time.Now().UnixNano())
    return rand.Float64() < dificulty
}


func callBackCatch(cfg *config,mapArea string,name string)error{
fmt.Println("Throwing a Pokeball at "+ name,"...")
    baseExperience, err := cfg.pokeapiClient.GetPokemonInfo(name)
   if err != nil{
      return err
    }
    // Check if the Pokemon is already caught
isCaughtAlready := false
for _, pokemon := range cfg.myPokemons {
    if name == pokemon {
        fmt.Println("already caught", name)
        isCaughtAlready = true
        break // Exit the loop if found
    }
}

// If not caught already, attempt to catch the Pokemon
if !isCaughtAlready {
    difficilty := 1 - (float64(baseExperience) / 350)
    isCaught := randomBool(difficilty)

    if isCaught {
        cfg.myPokemons = append(cfg.myPokemons, name)
        fmt.Println(name, "was caught!")
    } else {
        fmt.Println(name, "escaped!")
    }
}



    return nil
}
