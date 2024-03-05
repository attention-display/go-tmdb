
package guest_session

import (
	"github.com/attention-display/go-tmdb/conf"
	"github.com/attention-display/go-tmdb/repo/proxy"
)

type GuestSession struct {
	client *proxy.ProxyClient
}

func NewGuestSession(cfg *conf.Configuration) GuestSession {
	return GuestSession{
		client: proxy.NewProxyClient(cfg),
	}
}
