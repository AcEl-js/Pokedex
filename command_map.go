package main

import (
	"fmt"

	pokeapi "github.com/AcEl-js/Pokedex/pokesipa"
)

func callBackMap(){
     // get pokemon maps data
    pokeapiClient := pokeapi.NewClient()
    location,err := pokeapiClient.GetLocationData()

    if err != nil {
        fmt.Println("Error fetching location data: ",err)
        return
    }
    for _, result := range location.Results {
        fmt.Println("Location Name:", result.Name)
    }


}
