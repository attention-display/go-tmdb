package tv

import (
	"encoding/json"
	"fmt"
)

type DetailsReq struct {
	AppendToResponse string
	Language         string
}

func (r DetailsReq) Validate() error {
	return nil
}

func (r DetailsReq) getUrlPath() string {
	res := ""
	if r.AppendToResponse == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?append_to_response=%s", r.AppendToResponse)
	}

	if r.Language == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?language=%s", r.Language)
	}

	return res
}

type Detailscrew struct {
	Adult              bool
	Gender             int32
	Department         string
	Id                 int32
	KnownForDepartment string
	Name               string
	OriginalName       string
	Job                string
	Popularity         int32
	CreditId           string
	ProfilePath        string
}
type DetailsguestStars struct {
	CreditId           string
	Popularity         int32
	Character          string
	Adult              bool
	Name               string
	OriginalName       string
	Gender             int32
	KnownForDepartment string
	Order              int32
	ProfilePath        string
	Id                 int32
}
type Detailsepisodes struct {
	Id             int32
	Runtime        int32
	VoteAverage    int32
	Name           string
	Crew           []Detailscrew
	ShowId         int32
	AirDate        string
	GuestStars     []DetailsguestStars
	VoteCount      int32
	ProductionCode string
	Overview       string
	EpisodeNumber  int32
	SeasonNumber   int32
	StillPath      string
}

func UnmarshalDetailsResp(data []byte) (DetailsResp, error) {
	var r DetailsResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type DetailsResp struct {
	SeasonNumber int32
	VoteAverage  int32
	PosterPath   string
	Overview     string
	Episodes     []Detailsepisodes
	AirDate      string
	Id           int32
	_Id          string
	Name         string
}

func (a Tv) Details(series_id int32, season_number int32, req DetailsReq) (*DetailsResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/tv/%d/season/%d", a.client.Cfg.Host, series_id, season_number)

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
