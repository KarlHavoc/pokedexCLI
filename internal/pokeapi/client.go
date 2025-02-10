package pokeapi

import (
	"net/http"
	"time"

	"github.com/KarlHavoc/pokedexCLI/internal/pokecache"
)

// Client
type Client struct {
	httpclient http.Client
	cache      pokecache.Cache
}

// NewClient
func NewClient(timeout time.Duration, cacheInterval time.Duration) Client {
	return Client{
		httpclient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(cacheInterval),
	}
}
