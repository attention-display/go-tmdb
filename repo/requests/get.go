package requests

import (
	"github.com/attention-display/go-tmdb/consts"
)

func (r *Requests) Get(url string, headers map[string]string) (interface{}, error) {
	getResponse, err := SendRequest(consts.GET, url, headers, nil)
	if err != nil {
		return nil, err
	}
	return getResponse, nil
}
