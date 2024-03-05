package list

import (
	"encoding/json"
	"fmt"

	"github.com/attention-display/go-tmdb/entity/errors"
)

type CreateReq struct {
	SessionId string
}

func (r CreateReq) Validate() error {
	if r.SessionId == "" {
		return errors.InvalidParamsErr("SessionId is required.")
	}
	return nil
}

func (r CreateReq) getUrlPath() string {
	res := ""
	if r.SessionId == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?session_id=%s", r.SessionId)
	}

	return res
}

type CreateReqBody struct {
	RAW_BODY string
}

func (r CreateReqBody) Validate() error {
	if r.RAW_BODY == "" {
		return errors.InvalidParamsErr("RAW_BODY is required.")
	}
	return nil
}

func UnmarshalCreateResp(data []byte) (CreateResp, error) {
	var r CreateResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type CreateResp struct {
	Success       bool
	ListId        int32
	StatusCode    int32
	StatusMessage string
}

func (a List) Create(req CreateReq, reqBody CreateReqBody) (*CreateResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	if err := reqBody.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/list", a.client.Cfg.Host)

	resp, err := a.client.Urlib.Post(url, a.client.Cfg.DefaultHeaders, []byte(reqBody.RAW_BODY))
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalCreateResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
