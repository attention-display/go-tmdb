package keyword

import (
	"encoding/json"
	"fmt"
)

type MoviesReq struct {
	IncludeAdult bool
	Language     string
	Page         int32
}

func (r MoviesReq) Validate() error {
	return nil
}

func (r MoviesReq) getUrlPath() string {
	res := ""
	if r.Language == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?language=%s", r.Language)
	}

	return res
}

type Moviesresults struct {
	OriginalLanguage string
	Id               int32
	PosterPath       string
	Title            string
	Overview         string
	Adult            bool
	BackdropPath     string
	OriginalTitle    string
	ReleaseDate      string
	VoteCount        int32
	VoteAverage      int32
	Popularity       int32
	GenreIds         []int32
	Video            bool
}

func UnmarshalMoviesResp(data []byte) (MoviesResp, error) {
	var r MoviesResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type MoviesResp struct {
	TotalPages   int32
	TotalResults int32
	Id           int32
	Page         int32
	Results      []Moviesresults
}

func (a Keyword) Movies(keyword_id string, req MoviesReq) (*MoviesResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/keyword/%s/movies", a.client.Cfg.Host, keyword_id)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalMoviesResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
