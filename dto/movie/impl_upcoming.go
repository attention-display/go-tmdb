package movie

import (
	"encoding/json"
	"fmt"
)

type UpcomingReq struct {
	Page     int32
	Region   string
	Language string
}

func (r UpcomingReq) Validate() error {
	return nil
}

func (r UpcomingReq) getUrlPath() string {
	res := ""
	if r.Language == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?language=%s", r.Language)
	}

	if r.Region == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?region=%s", r.Region)
	}

	return res
}

type Upcomingdates struct {
	Maximum string
	Minimum string
}
type Upcomingresults struct {
	Popularity       int32
	PosterPath       string
	Id               int32
	Adult            bool
	OriginalTitle    string
	GenreIds         []int32
	BackdropPath     string
	Overview         string
	Title            string
	Video            bool
	VoteAverage      int32
	VoteCount        int32
	OriginalLanguage string
	ReleaseDate      string
}

func UnmarshalUpcomingResp(data []byte) (UpcomingResp, error) {
	var r UpcomingResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type UpcomingResp struct {
	Dates        Upcomingdates
	Page         int32
	Results      []Upcomingresults
	TotalPages   int32
	TotalResults int32
}

func (a Movie) Upcoming(req UpcomingReq) (*UpcomingResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/movie/upcoming", a.client.Cfg.Host)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalUpcomingResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
