package movie

import (
	"encoding/json"
	"fmt"
)

type MovieListReq struct {
	Page      int32
	StartDate string
	EndDate   string
}

func (r MovieListReq) Validate() error {
	return nil
}

func (r MovieListReq) getUrlPath() string {
	res := ""
	if r.StartDate == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?start_date=%s", r.StartDate)
	}

	if r.EndDate == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?end_date=%s", r.EndDate)
	}

	return res
}

type MovieListresults struct {
	Id    int32
	Adult bool
}

func UnmarshalMovieListResp(data []byte) (MovieListResp, error) {
	var r MovieListResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type MovieListResp struct {
	TotalResults int32
	Page         int32
	Results      []MovieListresults
	TotalPages   int32
}

func (a Movie) MovieList(req MovieListReq) (*MovieListResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/movie/changes", a.client.Cfg.Host)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalMovieListResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
