package movie

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

	if r.GuestSessionId == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?guest_session_id=%s", r.GuestSessionId)
	}

	return res
}

func (a Movie) DeleteRating(movie_id int32, req DeleteRatingReq) (*proto.CommonResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/movie/%d/rating", a.client.Cfg.Host, movie_id)

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
