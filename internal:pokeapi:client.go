package pokeapi

import (
	"net/http"
	"time"
)

// Client
type Client struct {
	httpclient http.Client
}

// NewClient
func NewClient(timeout time.Duration) Client {
	return Client{
		httpclient: http.Client{
			Timeout: timeout,
		},
	}
}
