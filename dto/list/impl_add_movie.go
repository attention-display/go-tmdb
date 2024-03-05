package list

import (
	"fmt"

	"github.com/attention-display/go-tmdb/entity/errors"
	"github.com/attention-display/go-tmdb/entity/proto"
)

type AddMovieReq struct {
	SessionId string
}

func (r AddMovieReq) Validate() error {
	if r.SessionId == "" {
		return errors.InvalidParamsErr("SessionId is required.")
	}
	return nil
}

func (r AddMovieReq) getUrlPath() string {
	res := ""
	if r.SessionId == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?session_id=%s", r.SessionId)
	}

	return res
}

type AddMovieReqBody struct {
	RAW_BODY string
}

func (r AddMovieReqBody) Validate() error {
	return nil
}

func (a List) AddMovie(list_id int32, req AddMovieReq, reqBody AddMovieReqBody) (*proto.CommonResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	if err := reqBody.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/list/%d/add_item", a.client.Cfg.Host, list_id)

	resp, err := a.client.Urlib.Post(url, a.client.Cfg.DefaultHeaders, nil)
	if err != nil {
		return nil, err
	}
	res, err := proto.UnmarshalCommonResponse(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
