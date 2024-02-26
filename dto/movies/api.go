package movies

import (
	"fmt"

	"github.com/attention-display/go-tmdb/repo/requests"
)

type Movies struct {
	client requests.Requests
}

func (m *Movies) Details() {
	fmt.Println("movies details")
}
