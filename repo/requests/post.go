package requests

import (
	"github.com/attention-display/go-tmdb/consts"
)

func (r *Requests) Post(url string, headers map[string]string, postBody []byte) (interface{}, error) {
	return SendRequest(consts.POST, url, headers, postBody)
}
