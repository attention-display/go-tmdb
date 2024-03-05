package configuration

import (
	"encoding/json"
	"fmt"
)

type Detailsimages struct {
	BaseUrl       string
	LogoSizes     []string
	PosterSizes   []string
	ProfileSizes  []string
	SecureBaseUrl string
	StillSizes    []string
	BackdropSizes []string
}

func UnmarshalDetailsResp(data []byte) (DetailsResp, error) {
	var r DetailsResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type DetailsResp struct {
	ChangeKeys []string
	Images     Detailsimages
}

func (a Configuration) Details() (*DetailsResp, error) {
	url := fmt.Sprintf("%s/3/configuration", a.client.Cfg.Host)

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
