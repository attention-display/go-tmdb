package account

import (
	"encoding/json"
	"fmt"
)

type WatchlistTVReq struct {
	Language  string
	Page      int32
	SessionId string
	SortBy    string
}

func (r WatchlistTVReq) Validate() error {
	return nil
}

func (r WatchlistTVReq) getUrlPath() string {
	res := ""
	if r.Language == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?language=%s", r.Language)
	}

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

	return res
}

type WatchlistTVresults struct {
	BackdropPath     string
	Adult            bool
	VoteCount        int32
	OriginalName     string
	GenreIds         []int32
	PosterPath       string
	FirstAirDate     string
	Name             string
	Overview         string
	VoteAverage      int32
	OriginalLanguage string
	OriginCountry    []string
	Popularity       int32
	Id               int32
}

func UnmarshalWatchlistTVResp(data []byte) (WatchlistTVResp, error) {
	var r WatchlistTVResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type WatchlistTVResp struct {
	TotalResults int32
	Page         int32
	Results      []WatchlistTVresults
	TotalPages   int32
}

func (a Account) WatchlistTV(account_id int32, req WatchlistTVReq) (*WatchlistTVResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/account/%d/watchlist/tv", a.client.Cfg.Host, account_id)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalWatchlistTVResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
