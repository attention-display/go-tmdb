package requests

import (
	"github.com/attention-display/go-tmdb/consts"
)

func (r *HttpClient) Get(url string, headers map[string]string) ([]byte, error) {
	return sendRequest(consts.GET, url, headers, nil)
}
