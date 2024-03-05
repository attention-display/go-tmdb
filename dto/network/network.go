
package network

import (
	"github.com/attention-display/go-tmdb/conf"
	"github.com/attention-display/go-tmdb/repo/proxy"
)

type Network struct {
	client *proxy.ProxyClient
}

func NewNetwork(cfg *conf.Configuration) Network {
	return Network{
		client: proxy.NewProxyClient(cfg),
	}
}
