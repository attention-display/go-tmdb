
package genre

import (
	"github.com/attention-display/go-tmdb/conf"
	"github.com/attention-display/go-tmdb/repo/proxy"
)

type Genre struct {
	client *proxy.ProxyClient
}

func NewGenre(cfg *conf.Configuration) Genre {
	return Genre{
		client: proxy.NewProxyClient(cfg),
	}
}
