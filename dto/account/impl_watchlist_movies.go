package account

import (
	"encoding/json"
	"fmt"
)

type WatchlistMoviesReq struct {
	SortBy    string
	Language  string
	Page      int32
	SessionId string
}

func (r WatchlistMoviesReq) Validate() error {
	return nil
}

func (r WatchlistMoviesReq) getUrlPath() string {
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

type WatchlistMoviesresults struct {
	BackdropPath     string
	OriginalLanguage string
	Overview         string
	PosterPath       string
	ReleaseDate      string
	Title            string
	Adult            bool
	Video            bool
	Popularity       int32
	VoteAverage      int32
	VoteCount        int32
	Id               int32
	GenreIds         []int32
	OriginalTitle    string
}

func UnmarshalWatchlistMoviesResp(data []byte) (WatchlistMoviesResp, error) {
	var r WatchlistMoviesResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type WatchlistMoviesResp struct {
	TotalPages   int32
	TotalResults int32
	Page         int32
	Results      []WatchlistMoviesresults
}

func (a Account) WatchlistMovies(account_id int32, req WatchlistMoviesReq) (*WatchlistMoviesResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/account/%d/watchlist/movies", a.client.Cfg.Host, account_id)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalWatchlistMoviesResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
