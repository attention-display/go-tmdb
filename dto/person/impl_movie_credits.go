package person

import (
	"encoding/json"
	"fmt"
)

type MovieCreditsReq struct {
	Language string
}

func (r MovieCreditsReq) Validate() error {
	return nil
}

func (r MovieCreditsReq) getUrlPath() string {
	res := ""
	if r.Language == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?language=%s", r.Language)
	}

	return res
}

type MovieCreditscast struct {
	Title            string
	Video            bool
	VoteCount        int32
	GenreIds         []int32
	Overview         string
	PosterPath       string
	Character        string
	Adult            bool
	CreditId         string
	Order            int32
	VoteAverage      int32
	OriginalLanguage string
	Popularity       int32
	ReleaseDate      string
	BackdropPath     string
	Id               int32
	OriginalTitle    string
}
type MovieCreditscrew struct {
	VoteCount        int32
	OriginalTitle    string
	Adult            bool
	Title            string
	Department       string
	GenreIds         []int32
	Video            bool
	PosterPath       string
	CreditId         string
	Id               int32
	Job              string
	VoteAverage      int32
	OriginalLanguage string
	ReleaseDate      string
	BackdropPath     string
	Overview         string
	Popularity       int32
}

func UnmarshalMovieCreditsResp(data []byte) (MovieCreditsResp, error) {
	var r MovieCreditsResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type MovieCreditsResp struct {
	Id   int32
	Cast []MovieCreditscast
	Crew []MovieCreditscrew
}

func (a Person) MovieCredits(person_id int32, req MovieCreditsReq) (*MovieCreditsResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/person/%d/movie_credits", a.client.Cfg.Host, person_id)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalMovieCreditsResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
