package tv

import (
	"encoding/json"
	"fmt"
)

type SimilarReq struct {
	Language string
	Page     int32
}

func (r SimilarReq) Validate() error {
	return nil
}

func (r SimilarReq) getUrlPath() string {
	res := ""
	if r.Language == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?language=%s", r.Language)
	}

	return res
}

type Similarresults struct {
	PosterPath       string
	FirstAirDate     string
	Name             string
	GenreIds         []int32
	VoteAverage      int32
	BackdropPath     string
	OriginCountry    []string
	OriginalLanguage string
	VoteCount        int32
	Adult            bool
	OriginalName     string
	Popularity       int32
	Overview         string
	Id               int32
}

func UnmarshalSimilarResp(data []byte) (SimilarResp, error) {
	var r SimilarResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type SimilarResp struct {
	Page         int32
	Results      []Similarresults
	TotalPages   int32
	TotalResults int32
}

func (a Tv) Similar(series_id string, req SimilarReq) (*SimilarResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/tv/%s/similar", a.client.Cfg.Host, series_id)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalSimilarResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
