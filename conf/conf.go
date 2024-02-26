package conf

type Config struct {
	// TMDb api key
	apiKey string
	// TMDb access token
	accessToken string
}

func NewConfig(opts ...Option) *Config {
	c := Config{}
	for _, f := range opts {
		f(&c)
	}
	return &c
}

// Option represents a functional option for configuring a TMDb instance.
type Option func(t *Config)

// WithApiKey is an option for setting the API key for TMDb.
func WithApiKey(apiKey string) Option {
	return func(t *Config) {
		t.apiKey = apiKey
	}
}

// WithAccessToken is an option for setting the access token for TMDb.
func WithAccessToken(accessToken string) Option {
	return func(t *Config) {
		t.accessToken = accessToken
	}
}

// SetApiKey sets the API key for the TMDb instance.
func (t *Config) SetApiKey(apiKey string) {
	t.apiKey = apiKey
}

// SetAccessToken sets the access token for the TMDb instance.
func (t *Config) SetAccessToken(accessToken string) {
	t.accessToken = accessToken
}
