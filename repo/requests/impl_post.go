package requests

import (
	"github.com/attention-display/go-tmdb/consts"
)

func (r *HttpClient) Post(url string, headers map[string]string, postBody []byte) ([]byte, error) {
	return sendRequest(consts.POST, url, headers, postBody)
}
