package proxy

import (
	"github.com/attention-display/go-tmdb/conf"
	"github.com/attention-display/go-tmdb/repo/requests"
)

type ProxyClient struct {
	Urlib requests.HttpClient
	Cfg   *conf.Configuration
}

func NewProxyClient(cfg *conf.Configuration) *ProxyClient {
	client := requests.NewHttpClient()
	return &ProxyClient{
		Urlib: client,
		Cfg:   cfg,
	}
}
