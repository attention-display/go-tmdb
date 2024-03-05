
package person

import (
	"github.com/attention-display/go-tmdb/conf"
	"github.com/attention-display/go-tmdb/repo/proxy"
)

type Person struct {
	client *proxy.ProxyClient
}

func NewPerson(cfg *conf.Configuration) Person {
	return Person{
		client: proxy.NewProxyClient(cfg),
	}
}
