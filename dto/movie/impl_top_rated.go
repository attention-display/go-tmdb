package movie

import (
	"encoding/json"
	"fmt"
)

type TopRatedReq struct {
	Language string
	Page     int32
	Region   string
}

func (r TopRatedReq) Validate() error {
	return nil
}

func (r TopRatedReq) getUrlPath() string {
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

type TopRatedresults struct {
	BackdropPath     string
	ReleaseDate      string
	Title            string
	OriginalLanguage string
	Adult            bool
	PosterPath       string
	OriginalTitle    string
	Video            bool
	VoteAverage      int32
	Id               int32
	Popularity       int32
	GenreIds         []int32
	Overview         string
	VoteCount        int32
}

func UnmarshalTopRatedResp(data []byte) (TopRatedResp, error) {
	var r TopRatedResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type TopRatedResp struct {
	Page         int32
	Results      []TopRatedresults
	TotalPages   int32
	TotalResults int32
}

func (a Movie) TopRated(req TopRatedReq) (*TopRatedResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/movie/top_rated", a.client.Cfg.Host)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalTopRatedResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
