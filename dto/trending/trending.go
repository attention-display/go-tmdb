
package trending

import (
	"github.com/attention-display/go-tmdb/conf"
	"github.com/attention-display/go-tmdb/repo/proxy"
)

type Trending struct {
	client *proxy.ProxyClient
}

func NewTrending(cfg *conf.Configuration) Trending {
	return Trending{
		client: proxy.NewProxyClient(cfg),
	}
}
