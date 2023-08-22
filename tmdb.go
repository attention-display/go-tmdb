package gotmdb

import (
	"github.com/attention-display/go-tmdb/repo/requests"
)

// The Movie Database
type TMDb struct {
	// TMDb api key
	apiKey string
	// TMDb access token
	accessToken string
	// http client
	requests.Requests
}

// New the movie db client
func NewTMDb(opts ...Option) *TMDb {
	var t TMDb
	t.Requests = requests.Requests{}
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
