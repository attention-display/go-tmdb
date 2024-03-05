package credit

import (
	"encoding/json"
	"fmt"
)

type Detailsperson struct {
	MediaType          string
	OriginalName       string
	KnownForDepartment string
	Name               string
	Popularity         int32
	Adult              bool
	Gender             int32
	ProfilePath        string
	Id                 int32
}
type Detailsseasons struct {
	PosterPath   string
	SeasonNumber int32
	ShowId       int32
	AirDate      string
	EpisodeCount int32
	Id           int32
	Name         string
	Overview     string
}
type Detailsmedia struct {
	Episodes         []interface{}
	BackdropPath     string
	Popularity       int32
	PosterPath       string
	OriginalLanguage string
	FirstAirDate     string
	Overview         string
	OriginalName     string
	VoteCount        int32
	Adult            bool
	OriginCountry    []string
	Seasons          []Detailsseasons
	MediaType        string
	Name             string
	VoteAverage      int32
	Character        string
	Id               int32
	GenreIds         []int32
}

func UnmarshalDetailsResp(data []byte) (DetailsResp, error) {
	var r DetailsResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type DetailsResp struct {
	CreditType string
	Department string
	Id         string
	Job        string
	Media      Detailsmedia
	MediaType  string
	Person     Detailsperson
}

func (a Credit) Details(credit_id string) (*DetailsResp, error) {
	url := fmt.Sprintf("%s/3/credit/%s", a.client.Cfg.Host, credit_id)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalDetailsResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
