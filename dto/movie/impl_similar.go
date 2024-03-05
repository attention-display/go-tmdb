package movie

import (
	"encoding/json"
	"fmt"
)

type SimilarReq struct {
	Language string
	Page     int32
}

func (r SimilarReq) Validate() error {
	return nil
}

func (r SimilarReq) getUrlPath() string {
	res := ""
	if r.Language == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?language=%s", r.Language)
	}

	return res
}

type Similarresults struct {
	BackdropPath     string
	OriginalTitle    string
	GenreIds         []int32
	ReleaseDate      string
	Video            bool
	Overview         string
	VoteCount        int32
	VoteAverage      int32
	PosterPath       string
	Adult            bool
	Popularity       int32
	OriginalLanguage string
	Title            string
	Id               int32
}

func UnmarshalSimilarResp(data []byte) (SimilarResp, error) {
	var r SimilarResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type SimilarResp struct {
	TotalResults int32
	Page         int32
	Results      []Similarresults
	TotalPages   int32
}

func (a Movie) Similar(movie_id int32, req SimilarReq) (*SimilarResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/movie/%d/similar", a.client.Cfg.Host, movie_id)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalSimilarResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
