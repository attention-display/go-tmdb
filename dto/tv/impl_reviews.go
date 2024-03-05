package tv

import (
	"encoding/json"
	"fmt"
)

type ReviewsReq struct {
	Language string
	Page     int32
}

func (r ReviewsReq) Validate() error {
	return nil
}

func (r ReviewsReq) getUrlPath() string {
	res := ""
	if r.Language == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?language=%s", r.Language)
	}

	return res
}

type ReviewsauthorDetails struct {
	Username   string
	AvatarPath string
	Name       string
	Rating     int32
}
type Reviewsresults struct {
	UpdatedAt     string
	Url           string
	Author        string
	AuthorDetails ReviewsauthorDetails
	Content       string
	CreatedAt     string
	Id            string
}

func UnmarshalReviewsResp(data []byte) (ReviewsResp, error) {
	var r ReviewsResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type ReviewsResp struct {
	Id           int32
	Page         int32
	Results      []Reviewsresults
	TotalPages   int32
	TotalResults int32
}

func (a Tv) Reviews(series_id int32, req ReviewsReq) (*ReviewsResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/tv/%d/reviews", a.client.Cfg.Host, series_id)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalReviewsResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
