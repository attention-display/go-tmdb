package tv

import (
	"encoding/json"
	"fmt"
)

type TopRatedReq struct {
	Language string
	Page     int32
}

func (r TopRatedReq) Validate() error {
	return nil
}

func (r TopRatedReq) getUrlPath() string {
	res := ""
	if r.Language == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?language=%s", r.Language)
	}

	return res
}

type TopRatedresults struct {
	VoteAverage      int32
	VoteCount        int32
	Overview         string
	GenreIds         []int32
	OriginalLanguage string
	Id               int32
	Name             string
	OriginCountry    []string
	Popularity       int32
	BackdropPath     string
	FirstAirDate     string
	OriginalName     string
	PosterPath       string
}

func UnmarshalTopRatedResp(data []byte) (TopRatedResp, error) {
	var r TopRatedResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type TopRatedResp struct {
	TotalPages   int32
	TotalResults int32
	Page         int32
	Results      []TopRatedresults
}

func (a Tv) TopRated(req TopRatedReq) (*TopRatedResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/tv/top_rated", a.client.Cfg.Host)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalTopRatedResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
