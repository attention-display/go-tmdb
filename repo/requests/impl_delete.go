package requests

import "github.com/attention-display/go-tmdb/consts"

func (r *httpClientImpl) Delete(url string, headers map[string]string) (interface{}, error) {
	return sendRequest(consts.DELETE, url, headers, nil)
}
