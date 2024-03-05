package search

import (
	"encoding/json"
	"fmt"

	"github.com/attention-display/go-tmdb/entity/errors"
)

type PersonReq struct {
	Language     string
	Page         int32
	Query        string
	IncludeAdult bool
}

func (r PersonReq) Validate() error {
	if r.Query == "" {
		return errors.InvalidParamsErr("Query is required.")
	}
	return nil
}

func (r PersonReq) getUrlPath() string {
	res := ""
	if r.Query == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?query=%s", r.Query)
	}

	if r.Language == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?language=%s", r.Language)
	}

	return res
}

type PersonknownFor struct {
	Id               int32
	Popularity       int32
	VoteCount        int32
	Adult            bool
	VoteAverage      int32
	OriginalTitle    string
	MediaType        string
	Video            bool
	OriginalLanguage string
	BackdropPath     string
	PosterPath       string
	Overview         string
	GenreIds         []int32
	Title            string
	ReleaseDate      string
}
type Personresults struct {
	Adult              bool
	Popularity         int32
	ProfilePath        string
	Name               string
	Gender             int32
	OriginalName       string
	Id                 int32
	KnownFor           []PersonknownFor
	KnownForDepartment string
}

func UnmarshalPersonResp(data []byte) (PersonResp, error) {
	var r PersonResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type PersonResp struct {
	Results      []Personresults
	TotalPages   int32
	TotalResults int32
	Page         int32
}

func (a Search) Person(req PersonReq) (*PersonResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/search/person", a.client.Cfg.Host)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalPersonResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
