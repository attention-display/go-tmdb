package list

import (
	"fmt"

	"github.com/attention-display/go-tmdb/entity/errors"
	"github.com/attention-display/go-tmdb/entity/proto"
)

type RemoveMovieReq struct {
	SessionId string
}

func (r RemoveMovieReq) Validate() error {
	if r.SessionId == "" {
		return errors.InvalidParamsErr("SessionId is required.")
	}
	return nil
}

func (r RemoveMovieReq) getUrlPath() string {
	res := ""
	if r.SessionId == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?session_id=%s", r.SessionId)
	}

	return res
}

type RemoveMovieReqBody struct {
	RAW_BODY string
}

func (r RemoveMovieReqBody) Validate() error {
	if r.RAW_BODY == "" {
		return errors.InvalidParamsErr("RAW_BODY is required.")
	}
	return nil
}

func (a List) RemoveMovie(list_id int32, req RemoveMovieReq, reqBody RemoveMovieReqBody) (*proto.CommonResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	if err := reqBody.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/list/%d/remove_item", a.client.Cfg.Host, list_id)

	resp, err := a.client.Urlib.Post(url, a.client.Cfg.DefaultHeaders, []byte(reqBody.RAW_BODY))
	if err != nil {
		return nil, err
	}
	res, err := proto.UnmarshalCommonResponse(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
