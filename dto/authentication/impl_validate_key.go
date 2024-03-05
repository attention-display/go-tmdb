package authentication

import (
	"encoding/json"
	"fmt"
)

func UnmarshalValidateKeyResp(data []byte) (ValidateKeyResp, error) {
	var r ValidateKeyResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type ValidateKeyResp struct {
	StatusCode    int32
	StatusMessage string
	Success       bool
}

func (a Authentication) ValidateKey() (*ValidateKeyResp, error) {
	url := fmt.Sprintf("%s/3/authentication", a.client.Cfg.Host)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalValidateKeyResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
