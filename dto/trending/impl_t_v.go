package trending

import (
	"encoding/json"
	"fmt"
)

type TVReq struct {
	Language string
}

func (r TVReq) Validate() error {
	return nil
}

func (r TVReq) getUrlPath() string {
	res := ""
	if r.Language == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?language=%s", r.Language)
	}

	return res
}

type TVresults struct {
	MediaType        string
	OriginalLanguage string
	VoteAverage      int32
	PosterPath       string
	Name             string
	VoteCount        int32
	OriginalName     string
	Adult            bool
	Id               int32
	OriginCountry    []string
	GenreIds         []int32
	Overview         string
	FirstAirDate     string
	Popularity       int32
	BackdropPath     string
}

func UnmarshalTVResp(data []byte) (TVResp, error) {
	var r TVResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type TVResp struct {
	TotalPages   int32
	TotalResults int32
	Page         int32
	Results      []TVresults
}

func (a Trending) TV(time_window string, req TVReq) (*TVResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/trending/tv/%s", a.client.Cfg.Host, time_window)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalTVResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
