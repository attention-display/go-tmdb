package tv

import (
	"encoding/json"
	"fmt"
)

type OnTheAirReq struct {
	Language string
	Page     int32
	Timezone string
}

func (r OnTheAirReq) Validate() error {
	return nil
}

func (r OnTheAirReq) getUrlPath() string {
	res := ""
	if r.Timezone == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?timezone=%s", r.Timezone)
	}

	if r.Language == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?language=%s", r.Language)
	}

	return res
}

type OnTheAirresults struct {
	VoteCount        int32
	Overview         string
	FirstAirDate     string
	Name             string
	Popularity       int32
	BackdropPath     string
	PosterPath       string
	OriginalName     string
	GenreIds         []int32
	Id               int32
	OriginCountry    []string
	OriginalLanguage string
	VoteAverage      int32
}

func UnmarshalOnTheAirResp(data []byte) (OnTheAirResp, error) {
	var r OnTheAirResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type OnTheAirResp struct {
	TotalPages   int32
	TotalResults int32
	Page         int32
	Results      []OnTheAirresults
}

func (a Tv) OnTheAir(req OnTheAirReq) (*OnTheAirResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/tv/on_the_air", a.client.Cfg.Host)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalOnTheAirResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
