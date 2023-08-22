package requests

import (
	"github.com/attention-display/go-tmdb/consts"
)

func (r *Requests) Post(url string, headers map[string]string, postBody []byte) (interface{}, error) {
	postResponse, err := SendRequest(consts.POST, url, headers, postBody)
	if err != nil {
		return nil, err
	}
	return postResponse, nil
}
