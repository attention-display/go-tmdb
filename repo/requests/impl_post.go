package requests

import (
	"github.com/attention-display/go-tmdb/consts"
)

func (r *httpClientImpl) Post(url string, headers map[string]string, postBody []byte) (interface{}, error) {
	return sendRequest(consts.POST, url, headers, postBody)
}
