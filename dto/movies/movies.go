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

func (m *Movies) Details() {
	fmt.Println("movies details")
}
