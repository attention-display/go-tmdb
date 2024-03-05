
package review

import (
	"github.com/attention-display/go-tmdb/conf"
	"github.com/attention-display/go-tmdb/repo/proxy"
)

type Review struct {
	client *proxy.ProxyClient
}

func NewReview(cfg *conf.Configuration) Review {
	return Review{
		client: proxy.NewProxyClient(cfg),
	}
}
