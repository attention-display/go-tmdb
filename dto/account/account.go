package account

import (
	"fmt"

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

func (a Account) getUrlPath() string {
	return fmt.Sprintf("%s/%s", a.client.Cfg.GetUrl(), "account")
}
