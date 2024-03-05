package authentication

import (
	"encoding/json"
	"fmt"

	"github.com/attention-display/go-tmdb/entity/errors"
)

type CreateSessionfromv4tokenReqBody struct {
	RAW_BODY string
}

func (r CreateSessionfromv4tokenReqBody) Validate() error {
	if r.RAW_BODY == "" {
		return errors.InvalidParamsErr("RAW_BODY is required.")
	}
	return nil
}

func UnmarshalCreateSessionfromv4tokenResp(data []byte) (CreateSessionfromv4tokenResp, error) {
	var r CreateSessionfromv4tokenResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type CreateSessionfromv4tokenResp struct {
	SessionId string
	Success   bool
}

func (a Authentication) CreateSessionfromv4token(reqBody CreateSessionfromv4tokenReqBody) (*CreateSessionfromv4tokenResp, error) {
	if err := reqBody.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/authentication/session/convert/4", a.client.Cfg.Host)

	resp, err := a.client.Urlib.Post(url, a.client.Cfg.DefaultHeaders, []byte(reqBody.RAW_BODY))
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalCreateSessionfromv4tokenResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
