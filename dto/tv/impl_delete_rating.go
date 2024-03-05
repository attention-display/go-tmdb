package tv

import (
	"fmt"

	"github.com/attention-display/go-tmdb/entity/proto"
)

type DeleteRatingReq struct {
	ContentType    string
	GuestSessionId string
	SessionId      string
}

func (r DeleteRatingReq) Validate() error {
	return nil
}

func (r DeleteRatingReq) getUrlPath() string {
	res := ""
	if r.ContentType == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?Content-Type=%s", r.ContentType)
	}

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

	return res
}

func (a Tv) DeleteRating(series_id int32, season_number int32, episode_number int32, req DeleteRatingReq) (*proto.CommonResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/tv/%d/season/%d/episode/%d/rating", a.client.Cfg.Host, series_id, season_number, episode_number)

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
