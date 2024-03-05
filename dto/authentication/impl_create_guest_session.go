package authentication

import (
	"encoding/json"
	"fmt"
)

func UnmarshalCreateGuestSessionResp(data []byte) (CreateGuestSessionResp, error) {
	var r CreateGuestSessionResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type CreateGuestSessionResp struct {
	ExpiresAt      string
	GuestSessionId string
	Success        bool
}

func (a Authentication) CreateGuestSession() (*CreateGuestSessionResp, error) {
	url := fmt.Sprintf("%s/3/authentication/guest_session/new", a.client.Cfg.Host)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalCreateGuestSessionResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
