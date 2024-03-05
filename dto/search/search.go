
package search

import (
	"github.com/attention-display/go-tmdb/conf"
	"github.com/attention-display/go-tmdb/repo/proxy"
)

type Search struct {
	client *proxy.ProxyClient
}

func NewSearch(cfg *conf.Configuration) Search {
	return Search{
		client: proxy.NewProxyClient(cfg),
	}
}
