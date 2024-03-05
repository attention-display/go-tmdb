
package list

import (
	"github.com/attention-display/go-tmdb/conf"
	"github.com/attention-display/go-tmdb/repo/proxy"
)

type List struct {
	client *proxy.ProxyClient
}

func NewList(cfg *conf.Configuration) List {
	return List{
		client: proxy.NewProxyClient(cfg),
	}
}
