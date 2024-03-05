package person

import (
	"encoding/json"
	"fmt"
)

func UnmarshalLatestResp(data []byte) (LatestResp, error) {
	var r LatestResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type LatestResp struct {
	Gender      int32
	Id          int32
	Name        string
	Biography   string
	Adult       bool
	AlsoKnownAs []interface{}
	Popularity  int32
}

func (a Person) Latest() (*LatestResp, error) {
	url := fmt.Sprintf("%s/3/person/latest", a.client.Cfg.Host)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalLatestResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
