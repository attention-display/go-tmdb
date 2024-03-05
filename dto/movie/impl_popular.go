package movie

import (
	"encoding/json"
	"fmt"
)

type PopularReq struct {
	Language string
	Page     int32
	Region   string
}

func (r PopularReq) Validate() error {
	return nil
}

func (r PopularReq) getUrlPath() string {
	res := ""
	if r.Region == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?region=%s", r.Region)
	}

	if r.Language == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?language=%s", r.Language)
	}

	return res
}

type Popularresults struct {
	Title            string
	Id               int32
	PosterPath       string
	GenreIds         []int32
	BackdropPath     string
	ReleaseDate      string
	Adult            bool
	Overview         string
	OriginalTitle    string
	Popularity       int32
	VoteCount        int32
	Video            bool
	VoteAverage      int32
	OriginalLanguage string
}

func UnmarshalPopularResp(data []byte) (PopularResp, error) {
	var r PopularResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type PopularResp struct {
	TotalResults int32
	Page         int32
	Results      []Popularresults
	TotalPages   int32
}

func (a Movie) Popular(req PopularReq) (*PopularResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/movie/popular", a.client.Cfg.Host)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalPopularResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
