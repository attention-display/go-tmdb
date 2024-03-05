package tv

import (
	"encoding/json"
	"fmt"
)

type Keywordsresults struct {
	Id   int32
	Name string
}

func UnmarshalKeywordsResp(data []byte) (KeywordsResp, error) {
	var r KeywordsResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type KeywordsResp struct {
	Id      int32
	Results []Keywordsresults
}

func (a Tv) Keywords(series_id int32) (*KeywordsResp, error) {
	url := fmt.Sprintf("%s/3/tv/%d/keywords", a.client.Cfg.Host, series_id)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalKeywordsResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
