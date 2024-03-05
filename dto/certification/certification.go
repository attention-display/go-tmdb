
package certification

import (
	"github.com/attention-display/go-tmdb/conf"
	"github.com/attention-display/go-tmdb/repo/proxy"
)

type Certification struct {
	client *proxy.ProxyClient
}

func NewCertification(cfg *conf.Configuration) Certification {
	return Certification{
		client: proxy.NewProxyClient(cfg),
	}
}
