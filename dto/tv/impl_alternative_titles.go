package tv

import (
	"encoding/json"
	"fmt"
)

type AlternativeTitlesresults struct {
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
	Id      int32
	Results []AlternativeTitlesresults
}

func (a Tv) AlternativeTitles(series_id int32) (*AlternativeTitlesResp, error) {
	url := fmt.Sprintf("%s/3/tv/%d/alternative_titles", a.client.Cfg.Host, series_id)

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
