package genre

import (
	"encoding/json"
	"fmt"
)

type MovieListReq struct {
	Language string
}

func (r MovieListReq) Validate() error {
	return nil
}

func (r MovieListReq) getUrlPath() string {
	res := ""
	if r.Language == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?language=%s", r.Language)
	}

	return res
}

type MovieListgenres struct {
	Id   int32
	Name string
}

func UnmarshalMovieListResp(data []byte) (MovieListResp, error) {
	var r MovieListResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type MovieListResp struct {
	Genres []MovieListgenres
}

func (a Genre) MovieList(req MovieListReq) (*MovieListResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/genre/movie/list", a.client.Cfg.Host)

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
