
package credit

import (
	"github.com/attention-display/go-tmdb/conf"
	"github.com/attention-display/go-tmdb/repo/proxy"
)

type Credit struct {
	client *proxy.ProxyClient
}

func NewCredit(cfg *conf.Configuration) Credit {
	return Credit{
		client: proxy.NewProxyClient(cfg),
	}
}
