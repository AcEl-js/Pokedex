package pokeapi

import (
    "encoding/json"
    "fmt"
    "net/http"
)

func (c *Client) GetLocationData() (Location,error) {
    
    endpoint := "/location-area"
    fullUrl := baseUrl + endpoint
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
return location,nil

}
