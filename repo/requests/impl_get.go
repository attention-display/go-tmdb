package requests

import (
	"github.com/attention-display/go-tmdb/consts"
)

func (r *httpClientImpl) Get(url string, headers map[string]string) (interface{}, error) {
	return sendRequest(consts.GET, url, headers, nil)
}
