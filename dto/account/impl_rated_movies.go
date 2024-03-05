package account

import (
	"encoding/json"
	"fmt"
)

type RatedMoviesReq struct {
	Language  string
	Page      int32
	SessionId string
	SortBy    string
}

func (r RatedMoviesReq) Validate() error {
	return nil
}

func (r RatedMoviesReq) getUrlPath() string {
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

type RatedMoviesresults struct {
	GenreIds         []int32
	Id               int32
	OriginalTitle    string
	ReleaseDate      string
	VoteAverage      int32
	PosterPath       string
	Overview         string
	VoteCount        int32
	OriginalLanguage string
	Title            string
	Popularity       int32
	Adult            bool
	Rating           int32
	Video            bool
	BackdropPath     string
}

func UnmarshalRatedMoviesResp(data []byte) (RatedMoviesResp, error) {
	var r RatedMoviesResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type RatedMoviesResp struct {
	Page         int32
	Results      []RatedMoviesresults
	TotalPages   int32
	TotalResults int32
}

func (a Account) RatedMovies(account_id int32, req RatedMoviesReq) (*RatedMoviesResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/account/%d/rated/movies", a.client.Cfg.Host, account_id)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalRatedMoviesResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
