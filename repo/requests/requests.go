package requests

import (
	"bytes"
	"io"
	"net/http"
)

//go:generate mockgen -destination=mock/mock_http_client.go -package mock -source=requests.go
type HttpClient interface {
	Post(url string, headers map[string]string, postBody []byte) (interface{}, error)
	Get(url string, headers map[string]string) (interface{}, error)
	Delete(url string, headers map[string]string) (interface{}, error)
}

type httpClientImpl struct{}

func NewHttpClient() httpClientImpl {
	return httpClientImpl{}
}

func sendRequest(method, url string, headers map[string]string, body []byte) ([]byte, error) {
	request, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		request.Header.Set(key, value)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
