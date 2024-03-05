package authentication

import (
	"encoding/json"
	"fmt"

	"github.com/attention-display/go-tmdb/entity/errors"
)

type CreateSessionReqBody struct {
	RAW_BODY string
}

func (r CreateSessionReqBody) Validate() error {
	if r.RAW_BODY == "" {
		return errors.InvalidParamsErr("RAW_BODY is required.")
	}
	return nil
}

func UnmarshalCreateSessionResp(data []byte) (CreateSessionResp, error) {
	var r CreateSessionResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type CreateSessionResp struct {
	SessionId string
	Success   bool
}

func (a Authentication) CreateSession(reqBody CreateSessionReqBody) (*CreateSessionResp, error) {
	if err := reqBody.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/authentication/session/new", a.client.Cfg.Host)

	resp, err := a.client.Urlib.Post(url, a.client.Cfg.DefaultHeaders, []byte(reqBody.RAW_BODY))
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalCreateSessionResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
