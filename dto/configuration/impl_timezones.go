package configuration

import (
	"encoding/json"
	"fmt"
)

func UnmarshalTimezonesResp(data []byte) (TimezonesResp, error) {
	var r TimezonesResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type TimezonesResp struct {
	Iso_3166_1 string
	Zones      []string
}

func (a Configuration) Timezones() (*TimezonesResp, error) {
	url := fmt.Sprintf("%s/3/configuration/timezones", a.client.Cfg.Host)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalTimezonesResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
