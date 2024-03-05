
package company

import (
	"github.com/attention-display/go-tmdb/conf"
	"github.com/attention-display/go-tmdb/repo/proxy"
)

type Company struct {
	client *proxy.ProxyClient
}

func NewCompany(cfg *conf.Configuration) Company {
	return Company{
		client: proxy.NewProxyClient(cfg),
	}
}
