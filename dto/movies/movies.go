package movies

import (
	"fmt"

	"github.com/attention-display/go-tmdb/conf"
	"github.com/attention-display/go-tmdb/repo/proxy"
)

type Movies struct {
	client *proxy.ProxyClient
}

func NewMovies(cfg *conf.Configuration) Movies {
	return Movies{
		client: proxy.NewProxyClient(cfg),
	}
}

func (a Movies) getUrlPath() string {
	return fmt.Sprintf("%s/%s", a.client.Cfg.GetUrl(), "movie")
}
