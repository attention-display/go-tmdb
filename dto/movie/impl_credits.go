package movie

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

type Creditscast struct {
	CreditId           string
	OriginalName       string
	Popularity         int32
	Adult              bool
	CastId             int32
	Order              int32
	ProfilePath        string
	Gender             int32
	KnownForDepartment string
	Id                 int32
	Name               string
	Character          string
}
type Creditscrew struct {
	Adult              bool
	Gender             int32
	KnownForDepartment string
	Department         string
	Name               string
	Popularity         int32
	ProfilePath        string
	CreditId           string
	Id                 int32
	Job                string
	OriginalName       string
}

func UnmarshalCreditsResp(data []byte) (CreditsResp, error) {
	var r CreditsResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type CreditsResp struct {
	Cast []Creditscast
	Crew []Creditscrew
	Id   int32
}

func (a Movie) Credits(movie_id int32, req CreditsReq) (*CreditsResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/movie/%d/credits", a.client.Cfg.Host, movie_id)

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
