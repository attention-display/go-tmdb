package gotmdb

import (
	"github.com/attention-display/go-tmdb/conf"
	"github.com/attention-display/go-tmdb/dto"
)

type tmdbSvr struct {
	Cfg *conf.Configuration
	dto.Controller
}

func NewTMDbClient(c *conf.Configuration) *tmdbSvr {
	return &tmdbSvr{
		Cfg:        c,
		Controller: dto.NewController(c),
	}
}
