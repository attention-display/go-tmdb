package requests

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

//go:generate mockgen -destination=mock/mock_http_client.go -package mock -source=requests.go
type HttpClientInterface interface {
	Post(url string, headers map[string]string, postBody []byte) ([]byte, error)
	Get(url string, headers map[string]string) ([]byte, error)
	Delete(url string, headers map[string]string, postBody []byte) ([]byte, error)
}

type HttpClient struct{}

func NewHttpClient() HttpClient {
	return HttpClient{}
}

type errorResponse struct {
	Success       bool
	StatusCode    int
	StatusMessage string
}

func unmarshalErrResponse(data []byte) (errorResponse, error) {
	var r errorResponse
	err := json.Unmarshal(data, &r)
	return r, err
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

	if response.StatusCode != 200 {
		res, err := unmarshalErrResponse(responseBody)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(res.StatusMessage)
	}
	return responseBody, nil
}
