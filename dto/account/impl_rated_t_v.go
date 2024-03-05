package account

import (
	"encoding/json"
	"fmt"
)

type RatedTVReq struct {
	SessionId string
	SortBy    string
	Language  string
	Page      int32
}

func (r RatedTVReq) Validate() error {
	return nil
}

func (r RatedTVReq) getUrlPath() string {
	res := ""
	if r.SessionId == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?session_id=%s", r.SessionId)
	}

	if r.SortBy == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?sort_by=%s", r.SortBy)
	}

	if r.Language == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?language=%s", r.Language)
	}

	return res
}

type RatedTVresults struct {
	PosterPath       string
	OriginalName     string
	FirstAirDate     string
	Rating           int32
	Name             string
	OriginalLanguage string
	Popularity       int32
	Overview         string
	OriginCountry    []string
	BackdropPath     string
	GenreIds         []int32
	VoteAverage      int32
	VoteCount        int32
	Adult            bool
	Id               int32
}

func UnmarshalRatedTVResp(data []byte) (RatedTVResp, error) {
	var r RatedTVResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type RatedTVResp struct {
	Results      []RatedTVresults
	TotalPages   int32
	TotalResults int32
	Page         int32
}

func (a Account) RatedTV(account_id int32, req RatedTVReq) (*RatedTVResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/account/%d/rated/tv", a.client.Cfg.Host, account_id)

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
