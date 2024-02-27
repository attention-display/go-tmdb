package movies

import (
	"fmt"

	"github.com/attention-display/go-tmdb/repo/proxy"
)

type Movies struct {
	client *proxy.ProxyClient
}

func (m *Movies) Details() {
	fmt.Println("movies details")
}
