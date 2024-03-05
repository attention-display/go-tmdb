
package authentication

import (
	"github.com/attention-display/go-tmdb/conf"
	"github.com/attention-display/go-tmdb/repo/proxy"
)

type Authentication struct {
	client *proxy.ProxyClient
}

func NewAuthentication(cfg *conf.Configuration) Authentication {
	return Authentication{
		client: proxy.NewProxyClient(cfg),
	}
}
