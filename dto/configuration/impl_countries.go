package configuration

import (
	"encoding/json"
	"fmt"
)

type CountriesReq struct {
	Language string
}

func (r CountriesReq) Validate() error {
	return nil
}

func (r CountriesReq) getUrlPath() string {
	res := ""
	if r.Language == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?language=%s", r.Language)
	}

	return res
}

func UnmarshalCountriesResp(data []byte) (CountriesResp, error) {
	var r CountriesResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type CountriesResp struct {
	EnglishName string
	Iso_3166_1  string
	NativeName  string
}

func (a Configuration) Countries(req CountriesReq) (*CountriesResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/configuration/countries", a.client.Cfg.Host)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalCountriesResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
