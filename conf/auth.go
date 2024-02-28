package conf

type ApiConfig struct {
	// TMDb api key
	apiKey string
	// TMDb access token
	accessToken string
}

func NewApiConfig(opts ...Option) *ApiConfig {
	c := ApiConfig{}
	for _, f := range opts {
		f(&c)
	}
	return &c
}

// Option represents a functional option for configuring a TMDb instance.
type Option func(t *ApiConfig)

// WithApiKey is an option for setting the API key for TMDb.
func WithApiKey(apiKey string) Option {
	return func(t *ApiConfig) {
		t.apiKey = apiKey
	}
}

// WithAccessToken is an option for setting the access token for TMDb.
func WithAccessToken(accessToken string) Option {
	return func(t *ApiConfig) {
		t.accessToken = accessToken
	}
}

// SetApiKey sets the API key for the TMDb instance.
func (t *ApiConfig) SetApiKey(apiKey string) {
	t.apiKey = apiKey
}

// SetAccessToken sets the access token for the TMDb instance.
func (t *ApiConfig) SetAccessToken(accessToken string) {
	t.accessToken = accessToken
}
