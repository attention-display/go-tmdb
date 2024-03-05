
package collection

import (
	"github.com/attention-display/go-tmdb/conf"
	"github.com/attention-display/go-tmdb/repo/proxy"
)

type Collection struct {
	client *proxy.ProxyClient
}

func NewCollection(cfg *conf.Configuration) Collection {
	return Collection{
		client: proxy.NewProxyClient(cfg),
	}
}
