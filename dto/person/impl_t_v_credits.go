package person

import (
	"encoding/json"
	"fmt"
)

type TVCreditsReq struct {
	Language string
}

func (r TVCreditsReq) Validate() error {
	return nil
}

func (r TVCreditsReq) getUrlPath() string {
	res := ""
	if r.Language == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?language=%s", r.Language)
	}

	return res
}

type TVCreditscast struct {
	Popularity       int32
	Overview         string
	OriginalLanguage string
	CreditId         string
	FirstAirDate     string
	EpisodeCount     int32
	GenreIds         []int32
	Name             string
	OriginCountry    []string
	VoteCount        int32
	BackdropPath     string
	Adult            bool
	Id               int32
	OriginalName     string
	Character        string
	PosterPath       string
	VoteAverage      int32
}
type TVCreditscrew struct {
	Overview         string
	FirstAirDate     string
	GenreIds         []int32
	BackdropPath     string
	Department       string
	Name             string
	VoteAverage      int32
	Id               int32
	OriginCountry    []string
	OriginalName     string
	PosterPath       string
	OriginalLanguage string
	Job              string
	CreditId         string
	Adult            bool
	EpisodeCount     int32
	Popularity       int32
	VoteCount        int32
}

func UnmarshalTVCreditsResp(data []byte) (TVCreditsResp, error) {
	var r TVCreditsResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type TVCreditsResp struct {
	Cast []TVCreditscast
	Crew []TVCreditscrew
	Id   int32
}

func (a Person) TVCredits(person_id int32, req TVCreditsReq) (*TVCreditsResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/person/%d/tv_credits", a.client.Cfg.Host, person_id)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalTVCreditsResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
