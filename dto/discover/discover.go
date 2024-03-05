
package discover

import (
	"github.com/attention-display/go-tmdb/conf"
	"github.com/attention-display/go-tmdb/repo/proxy"
)

type Discover struct {
	client *proxy.ProxyClient
}

func NewDiscover(cfg *conf.Configuration) Discover {
	return Discover{
		client: proxy.NewProxyClient(cfg),
	}
}
