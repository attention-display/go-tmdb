package conf

import (
	"fmt"

	"github.com/attention-display/go-tmdb/consts"
)

type Configuration struct {
	Host           string
	Version        string
	DefaultHeaders map[string]string
	ApiConfig      *ApiConfig
}

func NewDefaultConf(auth *ApiConfig) *Configuration {
	cfg := &Configuration{
		Host:           consts.BASEURL,
		Version:        consts.Version,
		DefaultHeaders: make(map[string]string),
		ApiConfig:      auth,
	}
	cfg.SetHeaders("accept", "application/json")
	cfg.SetHeaders("Authorization", fmt.Sprintf("Bearer %s", auth.apiKey))
	return cfg
}

func (c *Configuration) SetHeaders(key, val string) {
	c.DefaultHeaders[key] = val
}
