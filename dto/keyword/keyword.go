
package keyword

import (
	"github.com/attention-display/go-tmdb/conf"
	"github.com/attention-display/go-tmdb/repo/proxy"
)

type Keyword struct {
	client *proxy.ProxyClient
}

func NewKeyword(cfg *conf.Configuration) Keyword {
	return Keyword{
		client: proxy.NewProxyClient(cfg),
	}
}
