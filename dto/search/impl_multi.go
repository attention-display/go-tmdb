package search

import (
	"encoding/json"
	"fmt"

	"github.com/attention-display/go-tmdb/entity/errors"
)

type MultiReq struct {
	Query        string
	IncludeAdult bool
	Language     string
	Page         int32
}

func (r MultiReq) Validate() error {
	if r.Query == "" {
		return errors.InvalidParamsErr("Query is required.")
	}
	return nil
}

func (r MultiReq) getUrlPath() string {
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

type Multiresults struct {
	BackdropPath     string
	MediaType        string
	PosterPath       string
	OriginalLanguage string
	Adult            bool
	ReleaseDate      string
	Popularity       int32
	Video            bool
	GenreIds         []int32
	VoteAverage      int32
	VoteCount        int32
	OriginalTitle    string
	Overview         string
	Id               int32
	Title            string
}

func UnmarshalMultiResp(data []byte) (MultiResp, error) {
	var r MultiResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type MultiResp struct {
	TotalResults int32
	Page         int32
	Results      []Multiresults
	TotalPages   int32
}

func (a Search) Multi(req MultiReq) (*MultiResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/search/multi", a.client.Cfg.Host)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalMultiResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
