package tv

import (
	"encoding/json"
	"fmt"
)

type Latestseasons struct {
	EpisodeCount int32
	Id           int32
	Name         string
	Overview     string
	SeasonNumber int32
}
type LatestlastEpisodeToAir struct {
	Name           string
	Overview       string
	ProductionCode string
	ShowId         int32
	EpisodeNumber  int32
	Id             int32
	VoteCount      int32
	AirDate        string
	SeasonNumber   int32
	VoteAverage    int32
}

func UnmarshalLatestResp(data []byte) (LatestResp, error) {
	var r LatestResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type LatestResp struct {
	Status              string
	Popularity          int32
	Languages           []interface{}
	Tagline             string
	Name                string
	ProductionCompanies []interface{}
	OriginCountry       []string
	Adult               bool
	FirstAirDate        string
	Networks            []interface{}
	VoteCount           int32
	Type                string
	ProductionCountries []interface{}
	NumberOfSeasons     int32
	Seasons             []Latestseasons
	Overview            string
	OriginalLanguage    string
	CreatedBy           []interface{}
	NumberOfEpisodes    int32
	SpokenLanguages     []interface{}
	LastAirDate         string
	LastEpisodeToAir    LatestlastEpisodeToAir
	InProduction        bool
	Id                  int32
	Genres              []interface{}
	Homepage            string
	VoteAverage         int32
	EpisodeRunTime      []interface{}
	OriginalName        string
}

func (a Tv) Latest() (*LatestResp, error) {
	url := fmt.Sprintf("%s/3/tv/latest", a.client.Cfg.Host)

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
