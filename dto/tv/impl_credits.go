package tv

import (
	"encoding/json"
	"fmt"
)

type CreditsReq struct {
	Language string
}

func (r CreditsReq) Validate() error {
	return nil
}

func (r CreditsReq) getUrlPath() string {
	res := ""
	if r.Language == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?language=%s", r.Language)
	}

	return res
}

type CreditsguestStars struct {
	Gender             int32
	Id                 int32
	Order              int32
	Popularity         int32
	Character          string
	OriginalName       string
	CreditId           string
	KnownForDepartment string
	Name               string
	ProfilePath        string
	Adult              bool
}
type Creditscast struct {
	Gender             int32
	KnownForDepartment string
	Order              int32
	Popularity         int32
	CreditId           string
	ProfilePath        string
	Character          string
	Name               string
	Id                 int32
	OriginalName       string
	Adult              bool
}
type Creditscrew struct {
	Id                 int32
	KnownForDepartment string
	Name               string
	ProfilePath        string
	Job                string
	Popularity         int32
	CreditId           string
	Adult              bool
	Gender             int32
	Department         string
	OriginalName       string
}

func UnmarshalCreditsResp(data []byte) (CreditsResp, error) {
	var r CreditsResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type CreditsResp struct {
	Crew       []Creditscrew
	GuestStars []CreditsguestStars
	Id         int32
	Cast       []Creditscast
}

func (a Tv) Credits(series_id int32, season_number int32, episode_number int32, req CreditsReq) (*CreditsResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/tv/%d/season/%d/episode/%d/credits", a.client.Cfg.Host, series_id, season_number, episode_number)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalCreditsResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
