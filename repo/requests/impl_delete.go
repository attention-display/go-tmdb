package requests

import "github.com/attention-display/go-tmdb/consts"

func (r *HttpClient) Delete(url string, headers map[string]string, postBody []byte) ([]byte, error) {
	return sendRequest(consts.DELETE, url, headers, postBody)
}
