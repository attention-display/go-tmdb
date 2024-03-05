package tv

import (
	"encoding/json"
	"fmt"
)

type AggregateCreditsReq struct {
	Language string
}

func (r AggregateCreditsReq) Validate() error {
	return nil
}

func (r AggregateCreditsReq) getUrlPath() string {
	res := ""
	if r.Language == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?language=%s", r.Language)
	}

	return res
}

type AggregateCreditsroles struct {
	Character    string
	CreditId     string
	EpisodeCount int32
}
type AggregateCreditscast struct {
	Order              int32
	OriginalName       string
	Popularity         int32
	ProfilePath        string
	Gender             int32
	Roles              []AggregateCreditsroles
	Id                 int32
	Name               string
	KnownForDepartment string
	TotalEpisodeCount  int32
	Adult              bool
}
type AggregateCreditsjobs struct {
	CreditId     string
	EpisodeCount int32
	Job          string
}
type AggregateCreditscrew struct {
	Adult              bool
	Jobs               []AggregateCreditsjobs
	KnownForDepartment string
	Department         string
	Id                 int32
	Popularity         int32
	OriginalName       string
	ProfilePath        string
	TotalEpisodeCount  int32
	Gender             int32
	Name               string
}

func UnmarshalAggregateCreditsResp(data []byte) (AggregateCreditsResp, error) {
	var r AggregateCreditsResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type AggregateCreditsResp struct {
	Cast []AggregateCreditscast
	Crew []AggregateCreditscrew
	Id   int32
}

func (a Tv) AggregateCredits(series_id int32, req AggregateCreditsReq) (*AggregateCreditsResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/tv/%d/aggregate_credits", a.client.Cfg.Host, series_id)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalAggregateCreditsResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
