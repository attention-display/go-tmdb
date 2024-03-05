package configuration

import (
	"encoding/json"
	"fmt"
)

func UnmarshalLanguagesResp(data []byte) (LanguagesResp, error) {
	var r LanguagesResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type LanguagesResp struct {
	EnglishName string
	Iso_639_1   string
	Name        string
}

func (a Configuration) Languages() (*LanguagesResp, error) {
	url := fmt.Sprintf("%s/3/configuration/languages", a.client.Cfg.Host)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalLanguagesResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
