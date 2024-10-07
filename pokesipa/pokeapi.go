package pokeapi

import (
	"net/http"
	"time"
)
type Location struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

const baseUrl = "https://pokeapi.co/api/v2"
type Client struct{

    httpClient http.Client
}

func NewClient() Client{
    return Client{
        httpClient: http.Client{
            Timeout: time.Minute,
        },
    }
}
