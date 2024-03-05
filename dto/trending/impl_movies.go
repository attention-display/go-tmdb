package trending

import (
	"encoding/json"
	"fmt"
)

type MoviesReq struct {
	Language string
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
	Overview         string
	Video            bool
	BackdropPath     string
	VoteAverage      int32
	VoteCount        int32
	Id               int32
	Title            string
	ReleaseDate      string
	OriginalLanguage string
	GenreIds         []int32
	PosterPath       string
	Adult            bool
	OriginalTitle    string
	Popularity       int32
	MediaType        string
}

func UnmarshalMoviesResp(data []byte) (MoviesResp, error) {
	var r MoviesResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type MoviesResp struct {
	TotalResults int32
	Page         int32
	Results      []Moviesresults
	TotalPages   int32
}

func (a Trending) Movies(time_window string, req MoviesReq) (*MoviesResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/trending/movie/%s", a.client.Cfg.Host, time_window)

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
