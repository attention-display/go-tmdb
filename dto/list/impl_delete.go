package list

import (
	"fmt"

	"github.com/attention-display/go-tmdb/entity/errors"
	"github.com/attention-display/go-tmdb/entity/proto"
)

type DeleteReq struct {
	SessionId string
}

func (r DeleteReq) Validate() error {
	if r.SessionId == "" {
		return errors.InvalidParamsErr("SessionId is required.")
	}
	return nil
}

func (r DeleteReq) getUrlPath() string {
	res := ""
	if r.SessionId == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?session_id=%s", r.SessionId)
	}

	return res
}

func (a List) Delete(list_id int32, req DeleteReq) (*proto.CommonResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/list/%d", a.client.Cfg.Host, list_id)

	resp, err := a.client.Urlib.Delete(url, a.client.Cfg.DefaultHeaders, nil)
	if err != nil {
		return nil, err
	}
	res, err := proto.UnmarshalCommonResponse(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
