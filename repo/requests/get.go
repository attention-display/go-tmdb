package requests

import (
	"github.com/attention-display/go-tmdb/consts"
)

func (r *Requests) Get(url string, headers map[string]string) (interface{}, error) {
	return SendRequest(consts.GET, url, headers, nil)
}
