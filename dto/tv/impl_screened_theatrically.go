package tv

import (
	"encoding/json"
	"fmt"
)

type ScreenedTheatricallyresults struct {
	Id            int32
	SeasonNumber  int32
	EpisodeNumber int32
}

func UnmarshalScreenedTheatricallyResp(data []byte) (ScreenedTheatricallyResp, error) {
	var r ScreenedTheatricallyResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type ScreenedTheatricallyResp struct {
	Id      int32
	Results []ScreenedTheatricallyresults
}

func (a Tv) ScreenedTheatrically(series_id int32) (*ScreenedTheatricallyResp, error) {
	url := fmt.Sprintf("%s/3/tv/%d/screened_theatrically", a.client.Cfg.Host, series_id)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalScreenedTheatricallyResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
