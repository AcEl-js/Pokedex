package main

import (
	"fmt"

)

func callBackMap(cfg *config) error{
     // get pokemon maps data
    location,err := cfg.pokeapiClient.GetLocationData(cfg.nextLocationUrl)

    if err != nil {
        fmt.Println("Error fetching location data: ",err)
        return err
    }
    for _, result := range location.Results {
        fmt.Println("Location Name:", result.Name)
    }
    cfg.nextLocationUrl = location.Next
    cfg.prevousLocationUrl = location.Previous
    return nil


}
