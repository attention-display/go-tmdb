package search

import (
	"encoding/json"
	"fmt"

	"github.com/attention-display/go-tmdb/entity/errors"
)

type TVReq struct {
	Query            string
	FirstAirDateYear int32
	IncludeAdult     bool
	Language         string
	Page             int32
	Year             int32
}

func (r TVReq) Validate() error {
	if r.Query == "" {
		return errors.InvalidParamsErr("Query is required.")
	}
	return nil
}

func (r TVReq) getUrlPath() string {
	res := ""
	if r.Language == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?language=%s", r.Language)
	}

	if r.Query == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?query=%s", r.Query)
	}

	return res
}

type TVresults struct {
	VoteCount        int32
	VoteAverage      int32
	FirstAirDate     string
	Adult            bool
	Overview         string
	BackdropPath     string
	Id               int32
	Name             string
	Popularity       int32
	OriginalLanguage string
	OriginalName     string
	PosterPath       string
	OriginCountry    []string
	GenreIds         []int32
}

func UnmarshalTVResp(data []byte) (TVResp, error) {
	var r TVResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type TVResp struct {
	Page         int32
	Results      []TVresults
	TotalPages   int32
	TotalResults int32
}

func (a Search) TV(req TVReq) (*TVResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/search/tv", a.client.Cfg.Host)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalTVResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
