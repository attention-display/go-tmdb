package authentication

import (
	"encoding/json"
	"fmt"

	"github.com/attention-display/go-tmdb/entity/errors"
)

type DeleteSessionReqBody struct {
	RAW_BODY string
}

func (r DeleteSessionReqBody) Validate() error {
	if r.RAW_BODY == "" {
		return errors.InvalidParamsErr("RAW_BODY is required.")
	}
	return nil
}

func UnmarshalDeleteSessionResp(data []byte) (DeleteSessionResp, error) {
	var r DeleteSessionResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type DeleteSessionResp struct {
	Success bool
}

func (a Authentication) DeleteSession(reqBody DeleteSessionReqBody) (*DeleteSessionResp, error) {
	if err := reqBody.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/authentication/session", a.client.Cfg.Host)

	resp, err := a.client.Urlib.Delete(url, a.client.Cfg.DefaultHeaders, []byte(reqBody.RAW_BODY))
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalDeleteSessionResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
