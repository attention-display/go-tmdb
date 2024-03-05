package authentication

import (
	"encoding/json"
	"fmt"

	"github.com/attention-display/go-tmdb/entity/errors"
)

type CreateSessionwithloginReqBody struct {
	RAW_BODY string
}

func (r CreateSessionwithloginReqBody) Validate() error {
	if r.RAW_BODY == "" {
		return errors.InvalidParamsErr("RAW_BODY is required.")
	}
	return nil
}

func UnmarshalCreateSessionwithloginResp(data []byte) (CreateSessionwithloginResp, error) {
	var r CreateSessionwithloginResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type CreateSessionwithloginResp struct {
	ExpiresAt    string
	RequestToken string
	Success      bool
}

func (a Authentication) CreateSessionwithlogin(reqBody CreateSessionwithloginReqBody) (*CreateSessionwithloginResp, error) {
	if err := reqBody.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/authentication/token/validate_with_login", a.client.Cfg.Host)

	resp, err := a.client.Urlib.Post(url, a.client.Cfg.DefaultHeaders, []byte(reqBody.RAW_BODY))
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalCreateSessionwithloginResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
