package proxy

import (
	"github.com/attention-display/go-tmdb/conf"
	"github.com/attention-display/go-tmdb/repo/requests"
)

type proxyClient struct {
	cfg  *conf.Config
	http requests.Requests
}
