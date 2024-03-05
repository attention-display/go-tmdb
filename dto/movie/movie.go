
package movie

import (
	"github.com/attention-display/go-tmdb/conf"
	"github.com/attention-display/go-tmdb/repo/proxy"
)

type Movie struct {
	client *proxy.ProxyClient
}

func NewMovie(cfg *conf.Configuration) Movie {
	return Movie{
		client: proxy.NewProxyClient(cfg),
	}
}
