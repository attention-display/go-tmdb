
package tv

import (
	"github.com/attention-display/go-tmdb/conf"
	"github.com/attention-display/go-tmdb/repo/proxy"
)

type Tv struct {
	client *proxy.ProxyClient
}

func NewTv(cfg *conf.Configuration) Tv {
	return Tv{
		client: proxy.NewProxyClient(cfg),
	}
}
