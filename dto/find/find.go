
package find

import (
	"github.com/attention-display/go-tmdb/conf"
	"github.com/attention-display/go-tmdb/repo/proxy"
)

type Find struct {
	client *proxy.ProxyClient
}

func NewFind(cfg *conf.Configuration) Find {
	return Find{
		client: proxy.NewProxyClient(cfg),
	}
}
