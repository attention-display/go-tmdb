package gotmdb

import (
	"fmt"

	"github.com/attention-display/go-tmdb/conf"
	"github.com/attention-display/go-tmdb/dto"
)

type tmdbSvr struct {
	cfg *conf.Configuration
	dto.Controller
}

func NewTMDbClient(c *conf.Configuration) *tmdbSvr {
	return &tmdbSvr{
		cfg:        c,
		Controller: dto.NewController(c),
	}
}

func (t *tmdbSvr) UpdateApikey(apikey string) {
	t.cfg.ApiConfig.SetApiKey(apikey)
	t.cfg.SetHeaders("Authorization", fmt.Sprintf("Bearer %s", apikey))
}
