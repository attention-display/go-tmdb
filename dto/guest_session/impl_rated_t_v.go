package guest_session

import (
	"encoding/json"
	"fmt"
)

type RatedTVReq struct {
	SortBy   string
	Language string
	Page     int32
}

func (r RatedTVReq) Validate() error {
	return nil
}

func (r RatedTVReq) getUrlPath() string {
	res := ""
	if r.Language == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?language=%s", r.Language)
	}

	if r.SortBy == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?sort_by=%s", r.SortBy)
	}

	return res
}

type RatedTVresults struct {
	Overview         string
	BackdropPath     string
	PosterPath       string
	OriginalName     string
	VoteAverage      int32
	Adult            bool
	OriginalLanguage string
	Rating           int32
	VoteCount        int32
	GenreIds         []int32
	Popularity       int32
	Id               int32
	OriginCountry    []string
	FirstAirDate     string
	Name             string
}

func UnmarshalRatedTVResp(data []byte) (RatedTVResp, error) {
	var r RatedTVResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type RatedTVResp struct {
	Page         int32
	Results      []RatedTVresults
	TotalPages   int32
	TotalResults int32
}

func (a GuestSession) RatedTV(guest_session_id string, req RatedTVReq) (*RatedTVResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/guest_session/%s/rated/tv", a.client.Cfg.Host, guest_session_id)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalRatedTVResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
