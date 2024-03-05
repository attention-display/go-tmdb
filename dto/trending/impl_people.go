package trending

import (
	"encoding/json"
	"fmt"
)

type PeopleReq struct {
	Language string
}

func (r PeopleReq) Validate() error {
	return nil
}

func (r PeopleReq) getUrlPath() string {
	res := ""
	if r.Language == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?language=%s", r.Language)
	}

	return res
}

type PeopleknownFor struct {
	PosterPath       string
	VoteCount        int32
	Adult            bool
	Id               int32
	Popularity       int32
	ReleaseDate      string
	Title            string
	MediaType        string
	Overview         string
	BackdropPath     string
	GenreIds         []int32
	OriginalLanguage string
	OriginalTitle    string
	Video            bool
	VoteAverage      int32
}
type Peopleresults struct {
	Gender             int32
	Name               string
	Adult              bool
	MediaType          string
	OriginalName       string
	Popularity         int32
	Id                 int32
	KnownForDepartment string
	ProfilePath        string
	KnownFor           []PeopleknownFor
}

func UnmarshalPeopleResp(data []byte) (PeopleResp, error) {
	var r PeopleResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type PeopleResp struct {
	Results      []Peopleresults
	TotalPages   int32
	TotalResults int32
	Page         int32
}

func (a Trending) People(time_window string, req PeopleReq) (*PeopleResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/trending/person/%s", a.client.Cfg.Host, time_window)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalPeopleResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
