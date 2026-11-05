package pokeapi

import (
	"net/http"
	"time"
	"github.com/emorales1534/pokedex/internal/pokecache"
)

//Client
type Client struct {
	httpClient http.Client
	cache pokecache.Cache
}

// New Client
func NewClient (timeout, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(cacheInterval),
	}
}
