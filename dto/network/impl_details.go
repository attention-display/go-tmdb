package network

import (
	"encoding/json"
	"fmt"
)

func UnmarshalDetailsResp(data []byte) (DetailsResp, error) {
	var r DetailsResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type DetailsResp struct {
	OriginCountry string
	Headquarters  string
	Homepage      string
	Id            int32
	LogoPath      string
	Name          string
}

func (a Network) Details(network_id int32) (*DetailsResp, error) {
	url := fmt.Sprintf("%s/3/network/%d", a.client.Cfg.Host, network_id)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalDetailsResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
