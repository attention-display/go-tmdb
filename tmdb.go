package gotmdb

import (
	"fmt"

	"github.com/attention-display/go-tmdb/consts"
)

// The Movie Database
type TMDb struct {
	// TMDb api key
	apiKey string
	// TMDb access token
	accessToken string
}

// New the movie db client
func NewTMDb(opts ...Option) *TMDb {
	var t TMDb
	for _, opt := range opts {
		opt(&t)
	}
	return &t
}

// Option represents a functional option for configuring a TMDb instance.
type Option func(t *TMDb)

// WithApiKey is an option for setting the API key for TMDb.
func WithApiKey(apiKey string) Option {
	return func(t *TMDb) {
		t.apiKey = apiKey
	}
}

// WithAccessToken is an option for setting the access token for TMDb.
func WithAccessToken(accessToken string) Option {
	return func(t *TMDb) {
		t.accessToken = accessToken
	}
}

// SetApiKey sets the API key for the TMDb instance.
func (t *TMDb) SetApiKey(apiKey string) {
	t.apiKey = apiKey
}

// SetAccessToken sets the access token for the TMDb instance.
func (t *TMDb) SetAccessToken(accessToken string) {
	t.accessToken = accessToken
}

// getUrlPath constructs the full URL path for the specified endpoint by combining the base URL,
// API version, and the given path.
func (t *TMDb) getUrlPath(path string) string {
	url := fmt.Sprintf("%s/%s/%s", consts.BASEURL, consts.Version, path)
	if t.apiKey != "" {
		url = fmt.Sprintf("%s?api_key=%s", url, t.apiKey)
	}
	return url
}

// getHeaders returns the headers to be included in the HTTP request.
func (t *TMDb) getHeaders() map[string]string {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	if t.accessToken != "" {
		headers["Authorization"] = fmt.Sprintf("Bearer %s", t.accessToken)
	}
	return headers
}
