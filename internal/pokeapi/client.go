package pokeapi

import (
	"net/http"
	"time"

	"github.com/vaha-b/go-cli-pokedex/internal/pokecache"
)

// Client -
type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

// NewClient -
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
