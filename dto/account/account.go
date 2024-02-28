package account

import (
	"github.com/attention-display/go-tmdb/conf"
	"github.com/attention-display/go-tmdb/repo/proxy"
)

type Account struct {
	client *proxy.ProxyClient
}

func NewAccount(cfg *conf.Configuration) Account {
	return Account{
		client: proxy.NewProxyClient(cfg),
	}
}
