
package watch

import (
	"github.com/attention-display/go-tmdb/conf"
	"github.com/attention-display/go-tmdb/repo/proxy"
)

type Watch struct {
	client *proxy.ProxyClient
}

func NewWatch(cfg *conf.Configuration) Watch {
	return Watch{
		client: proxy.NewProxyClient(cfg),
	}
}
