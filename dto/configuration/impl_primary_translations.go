package configuration

import "fmt"

func (a Configuration) PrimaryTranslations() (interface{}, error) {
	url := fmt.Sprintf("%s/3/configuration/primary_translations", a.client.Cfg.Host)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
