package review

import (
	"encoding/json"
	"fmt"
)

type DetailsauthorDetails struct {
	AvatarPath string
	Name       string
	Rating     int32
	Username   string
}

func UnmarshalDetailsResp(data []byte) (DetailsResp, error) {
	var r DetailsResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type DetailsResp struct {
	Url           string
	AuthorDetails DetailsauthorDetails
	CreatedAt     string
	MediaId       int32
	MediaTitle    string
	Author        string
	Content       string
	Iso_639_1     string
	MediaType     string
	UpdatedAt     string
	Id            string
}

func (a Review) Details(review_id string) (*DetailsResp, error) {
	url := fmt.Sprintf("%s/3/review/%s", a.client.Cfg.Host, review_id)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalDetailsResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
