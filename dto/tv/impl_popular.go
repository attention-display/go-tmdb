package tv

import (
	"encoding/json"
	"fmt"
)

type PopularReq struct {
	Language string
	Page     int32
}

func (r PopularReq) Validate() error {
	return nil
}

func (r PopularReq) getUrlPath() string {
	res := ""
	if r.Language == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?language=%s", r.Language)
	}

	return res
}

type Popularresults struct {
	GenreIds         []int32
	OriginalLanguage string
	VoteAverage      int32
	FirstAirDate     string
	OriginalName     string
	Popularity       int32
	VoteCount        int32
	Id               int32
	Name             string
	PosterPath       string
	BackdropPath     string
	OriginCountry    []string
	Overview         string
}

func UnmarshalPopularResp(data []byte) (PopularResp, error) {
	var r PopularResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type PopularResp struct {
	Page         int32
	Results      []Popularresults
	TotalPages   int32
	TotalResults int32
}

func (a Tv) Popular(req PopularReq) (*PopularResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/tv/popular", a.client.Cfg.Host)

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
