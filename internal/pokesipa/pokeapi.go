package pokeapi

import (
	"net/http"
	"time"
    	"github.com/AcEl-js/Pokedex/internal/pokecache"
)
type Location struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

const baseUrl = "https://pokeapi.co/api/v2"
type Client struct{

    httpClient http.Client
    cache      *pokecache.Cache
}

func NewClient(cleanupTick time.Duration) Client{
    return Client{
        httpClient: http.Client{
            Timeout: time.Minute,
        },
          cache: pokecache.NewCache(cleanupTick),
    }
}
