package person

import (
	"encoding/json"
	"fmt"
)

type PopularReq struct {
	Language string
	Page     int32
}

func (r PopularReq) Validate() error {
	return nil
}

func (r PopularReq) getUrlPath() string {
	res := ""
	if r.Language == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?language=%s", r.Language)
	}

	return res
}

type PopularknownFor struct {
	OriginalLanguage string
	Overview         string
	Id               int32
	PosterPath       string
	MediaType        string
	Adult            bool
	Title            string
	GenreIds         []int32
	VoteAverage      int32
	BackdropPath     string
	ReleaseDate      string
	Video            bool
	OriginalTitle    string
	VoteCount        int32
}
type Popularresults struct {
	ProfilePath        string
	Adult              bool
	Gender             int32
	Id                 int32
	KnownFor           []PopularknownFor
	KnownForDepartment string
	Name               string
	Popularity         int32
}

func UnmarshalPopularResp(data []byte) (PopularResp, error) {
	var r PopularResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type PopularResp struct {
	Page         int32
	Results      []Popularresults
	TotalPages   int32
	TotalResults int32
}

func (a Person) Popular(req PopularReq) (*PopularResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/person/popular", a.client.Cfg.Host)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalPopularResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
