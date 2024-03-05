package authentication

import (
	"encoding/json"
	"fmt"
)

func UnmarshalCreateRequestTokenResp(data []byte) (CreateRequestTokenResp, error) {
	var r CreateRequestTokenResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type CreateRequestTokenResp struct {
	Success      bool
	ExpiresAt    string
	RequestToken string
}

func (a Authentication) CreateRequestToken() (*CreateRequestTokenResp, error) {
	url := fmt.Sprintf("%s/3/authentication/token/new", a.client.Cfg.Host)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalCreateRequestTokenResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
