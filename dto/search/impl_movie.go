package search

import (
	"encoding/json"
	"fmt"

	"github.com/attention-display/go-tmdb/entity/errors"
)

type MovieReq struct {
	Language           string
	PrimaryReleaseYear string
	Page               int32
	Region             string
	Year               string
	Query              string
	IncludeAdult       bool
}

func (r MovieReq) Validate() error {
	if r.Query == "" {
		return errors.InvalidParamsErr("Query is required.")
	}
	return nil
}

func (r MovieReq) getUrlPath() string {
	res := ""
	if r.Region == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?region=%s", r.Region)
	}

	if r.Year == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?year=%s", r.Year)
	}

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

	if r.PrimaryReleaseYear == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?primary_release_year=%s", r.PrimaryReleaseYear)
	}

	return res
}

type Movieresults struct {
	VoteAverage      int32
	VoteCount        int32
	Overview         string
	Title            string
	OriginalTitle    string
	GenreIds         []int32
	Id               int32
	Adult            bool
	ReleaseDate      string
	PosterPath       string
	Video            bool
	Popularity       int32
	OriginalLanguage string
	BackdropPath     string
}

func UnmarshalMovieResp(data []byte) (MovieResp, error) {
	var r MovieResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type MovieResp struct {
	Page         int32
	Results      []Movieresults
	TotalPages   int32
	TotalResults int32
}

func (a Search) Movie(req MovieReq) (*MovieResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/search/movie", a.client.Cfg.Host)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalMovieResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
