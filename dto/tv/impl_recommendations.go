package tv

import (
	"encoding/json"
	"fmt"
)

type RecommendationsReq struct {
	Language string
	Page     int32
}

func (r RecommendationsReq) Validate() error {
	return nil
}

func (r RecommendationsReq) getUrlPath() string {
	res := ""
	if r.Language == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?language=%s", r.Language)
	}

	return res
}

type Recommendationsresults struct {
	BackdropPath     string
	FirstAirDate     string
	Name             string
	OriginCountry    []string
	PosterPath       string
	Overview         string
	MediaType        string
	Adult            bool
	Popularity       int32
	GenreIds         []int32
	OriginalLanguage string
	Id               int32
	OriginalName     string
	VoteCount        int32
	VoteAverage      int32
}

func UnmarshalRecommendationsResp(data []byte) (RecommendationsResp, error) {
	var r RecommendationsResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type RecommendationsResp struct {
	TotalResults int32
	Page         int32
	Results      []Recommendationsresults
	TotalPages   int32
}

func (a Tv) Recommendations(series_id int32, req RecommendationsReq) (*RecommendationsResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/tv/%d/recommendations", a.client.Cfg.Host, series_id)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalRecommendationsResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
