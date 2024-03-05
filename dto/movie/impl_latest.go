package movie

import (
	"encoding/json"
	"fmt"
)

func UnmarshalLatestResp(data []byte) (LatestResp, error) {
	var r LatestResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type LatestResp struct {
	Status              string
	Genres              []interface{}
	Popularity          int32
	Title               string
	OriginalLanguage    string
	Homepage            string
	ProductionCompanies []interface{}
	SpokenLanguages     []interface{}
	Tagline             string
	Overview            string
	OriginalTitle       string
	VoteCount           int32
	ProductionCountries []interface{}
	Revenue             int32
	Runtime             int32
	VoteAverage         int32
	Id                  int32
	Adult               bool
	ReleaseDate         string
	Video               bool
	Budget              int32
}

func (a Movie) Latest() (*LatestResp, error) {
	url := fmt.Sprintf("%s/3/movie/latest", a.client.Cfg.Host)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalLatestResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
