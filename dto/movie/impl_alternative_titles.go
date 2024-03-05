package movie

import (
	"encoding/json"
	"fmt"
)

type AlternativeTitlesReq struct {
	Country string
}

func (r AlternativeTitlesReq) Validate() error {
	return nil
}

func (r AlternativeTitlesReq) getUrlPath() string {
	res := ""
	if r.Country == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?country=%s", r.Country)
	}

	return res
}

type AlternativeTitlestitles struct {
	Iso_3166_1 string
	Title      string
	Type       string
}

func UnmarshalAlternativeTitlesResp(data []byte) (AlternativeTitlesResp, error) {
	var r AlternativeTitlesResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type AlternativeTitlesResp struct {
	Id     int32
	Titles []AlternativeTitlestitles
}

func (a Movie) AlternativeTitles(movie_id int32, req AlternativeTitlesReq) (*AlternativeTitlesResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/movie/%d/alternative_titles", a.client.Cfg.Host, movie_id)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalAlternativeTitlesResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
