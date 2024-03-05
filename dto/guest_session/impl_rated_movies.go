package guest_session

import (
	"encoding/json"
	"fmt"
)

type RatedMoviesReq struct {
	Language string
	Page     int32
	SortBy   string
}

func (r RatedMoviesReq) Validate() error {
	return nil
}

func (r RatedMoviesReq) getUrlPath() string {
	res := ""
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

type RatedMoviesresults struct {
	GenreIds         []int32
	Rating           int32
	Popularity       int32
	Video            bool
	OriginalLanguage string
	ReleaseDate      string
	Adult            bool
	BackdropPath     string
	Id               int32
	VoteCount        int32
	Overview         string
	Title            string
	VoteAverage      int32
	OriginalTitle    string
	PosterPath       string
}

func UnmarshalRatedMoviesResp(data []byte) (RatedMoviesResp, error) {
	var r RatedMoviesResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type RatedMoviesResp struct {
	TotalPages   int32
	TotalResults int32
	Page         int32
	Results      []RatedMoviesresults
}

func (a GuestSession) RatedMovies(guest_session_id string, req RatedMoviesReq) (*RatedMoviesResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/guest_session/%s/rated/movies", a.client.Cfg.Host, guest_session_id)

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
