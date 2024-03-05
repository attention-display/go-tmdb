package tv

import (
	"fmt"

	"github.com/attention-display/go-tmdb/entity/errors"
	"github.com/attention-display/go-tmdb/entity/proto"
)

type AddRatingReq struct {
	GuestSessionId string
	SessionId      string
	ContentType    string
}

func (r AddRatingReq) Validate() error {
	if r.ContentType == "" {
		return errors.InvalidParamsErr("ContentType is required.")
	}
	return nil
}

func (r AddRatingReq) getUrlPath() string {
	res := ""
	if r.GuestSessionId == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?guest_session_id=%s", r.GuestSessionId)
	}

	if r.SessionId == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?session_id=%s", r.SessionId)
	}

	if r.ContentType == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?Content-Type=%s", r.ContentType)
	}

	return res
}

type AddRatingReqBody struct {
	RAW_BODY string
}

func (r AddRatingReqBody) Validate() error {
	if r.RAW_BODY == "" {
		return errors.InvalidParamsErr("RAW_BODY is required.")
	}
	return nil
}

func (a Tv) AddRating(series_id int32, season_number int32, episode_number int32, req AddRatingReq, reqBody AddRatingReqBody) (*proto.CommonResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	if err := reqBody.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/tv/%d/season/%d/episode/%d/rating", a.client.Cfg.Host, series_id, season_number, episode_number)

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
