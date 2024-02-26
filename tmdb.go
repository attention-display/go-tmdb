package gotmdb

import (
	"github.com/attention-display/go-tmdb/conf"
	"github.com/attention-display/go-tmdb/dto"
)

// The Movie Database
type tmdbClient struct {
	cfg *conf.Config
	dto.Controller
}

func NewTMDbClient(c *conf.Config) *tmdbClient {
	return &tmdbClient{
		cfg:        c,
		Controller: dto.Controller{},
	}
}
