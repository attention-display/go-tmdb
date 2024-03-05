package person

import (
	"encoding/json"
	"fmt"
)

type CombinedCreditsReq struct {
	Language string
}

func (r CombinedCreditsReq) Validate() error {
	return nil
}

func (r CombinedCreditsReq) getUrlPath() string {
	res := ""
	if r.Language == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?language=%s", r.Language)
	}

	return res
}

type CombinedCreditscast struct {
	OriginalLanguage string
	BackdropPath     string
	VoteCount        int32
	Popularity       int32
	Overview         string
	Character        string
	Adult            bool
	CreditId         string
	Id               int32
	PosterPath       string
	GenreIds         []int32
	Order            int32
	Title            string
	MediaType        string
	ReleaseDate      string
	Video            bool
	VoteAverage      int32
	OriginalTitle    string
}
type CombinedCreditscrew struct {
	VoteCount        int32
	Title            string
	GenreIds         []int32
	Id               int32
	Department       string
	VoteAverage      int32
	OriginalLanguage string
	MediaType        string
	Overview         string
	ReleaseDate      string
	Video            bool
	Popularity       int32
	Job              string
	BackdropPath     string
	OriginalTitle    string
	PosterPath       string
	Adult            bool
	CreditId         string
}

func UnmarshalCombinedCreditsResp(data []byte) (CombinedCreditsResp, error) {
	var r CombinedCreditsResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type CombinedCreditsResp struct {
	Cast []CombinedCreditscast
	Crew []CombinedCreditscrew
	Id   int32
}

func (a Person) CombinedCredits(person_id string, req CombinedCreditsReq) (*CombinedCreditsResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/person/%s/combined_credits", a.client.Cfg.Host, person_id)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalCombinedCreditsResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
