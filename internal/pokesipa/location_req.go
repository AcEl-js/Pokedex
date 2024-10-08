package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetLocationData(pageUrl *string) (Location,error) {
    
    endpoint := "/location-area"
    fullUrl := baseUrl + endpoint
    if pageUrl != nil {
        fullUrl = *pageUrl
    }

    if cachedData, found := c.cache.Get(fullUrl);found{
        var location Location
        fmt.Println("cache hit!")
    

    if err := json.Unmarshal(cachedData, &location); err != nil {
    			return Location{}, err
		}
		return location, nil
    }
fmt.Println("cache miss!")
	req, err := http.NewRequest("GET",fullUrl, nil)
    if err != nil {
        return Location{},err
    }
    res, err := c.httpClient.Do(req)
    if err != nil{
        fmt.Println("error making requst: ",err)
        return Location{},err

    }
    
    
    defer res.Body.Close()
    
    var location Location 
    
    if err := json.NewDecoder(res.Body).Decode(&location);err !=nil{
        return Location{},err

    }
    data, err := json.Marshal(location)
	if err != nil {
		return Location{}, err
	}
  
    c.cache.Add(fullUrl, data)
    return location,nil

}
