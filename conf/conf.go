package conf

import (
	"github.com/attention-display/go-tmdb/consts"
)

type Configuration struct {
	Host      string
	Version   string
	Headers   map[string]string
	ApiConfig ApiConfig
}

func NewDefaultConf(auth ApiConfig) *Configuration {
	return &Configuration{
		Host:      consts.BASEURL,
		Version:   consts.Version,
		Headers:   make(map[string]string),
		ApiConfig: auth,
	}
}

func (c *Configuration) SetHeaders(key, val string) {
	c.Headers[key] = val
}
