
package configuration

import (
	"github.com/attention-display/go-tmdb/conf"
	"github.com/attention-display/go-tmdb/repo/proxy"
)

type Configuration struct {
	client *proxy.ProxyClient
}

func NewConfiguration(cfg *conf.Configuration) Configuration {
	return Configuration{
		client: proxy.NewProxyClient(cfg),
	}
}
