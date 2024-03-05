package tv

import (
	"encoding/json"
	"fmt"
)

type AiringTodayReq struct {
	Language string
	Page     int32
	Timezone string
}

func (r AiringTodayReq) Validate() error {
	return nil
}

func (r AiringTodayReq) getUrlPath() string {
	res := ""
	if r.Language == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?language=%s", r.Language)
	}

	if r.Timezone == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?timezone=%s", r.Timezone)
	}

	return res
}

type AiringTodayresults struct {
	Name             string
	Popularity       int32
	VoteAverage      int32
	GenreIds         []int32
	OriginCountry    []string
	OriginalName     string
	PosterPath       string
	VoteCount        int32
	FirstAirDate     string
	Id               int32
	OriginalLanguage string
	Overview         string
	BackdropPath     string
}

func UnmarshalAiringTodayResp(data []byte) (AiringTodayResp, error) {
	var r AiringTodayResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type AiringTodayResp struct {
	TotalResults int32
	Page         int32
	Results      []AiringTodayresults
	TotalPages   int32
}

func (a Tv) AiringToday(req AiringTodayReq) (*AiringTodayResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/tv/airing_today", a.client.Cfg.Host)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalAiringTodayResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
