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
    difficilty := 1 - (float64(baseExperience)/350)
    isCaught := randomBool(difficilty)
    fmt.Println("random :",isCaught)
    fmt.Println("dificulty :",difficilty)
    fmt.Println("baseEx :",baseExperience)
    var myPokemons []string
    
    if isCaught{
        myPokemons = append(myPokemons, name)
        fmt.Println(name,"was caught!")

    }else{

        fmt.Println(name,"escaped!")
    }

    
return nil
}
    
