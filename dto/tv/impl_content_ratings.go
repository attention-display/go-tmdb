package tv

import (
	"encoding/json"
	"fmt"
)

type ContentRatingsresults struct {
	Iso_3166_1  string
	Rating      string
	Descriptors []interface{}
}

func UnmarshalContentRatingsResp(data []byte) (ContentRatingsResp, error) {
	var r ContentRatingsResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type ContentRatingsResp struct {
	Results []ContentRatingsresults
	Id      int32
}

func (a Tv) ContentRatings(series_id int32) (*ContentRatingsResp, error) {
	url := fmt.Sprintf("%s/3/tv/%d/content_ratings", a.client.Cfg.Host, series_id)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalContentRatingsResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
