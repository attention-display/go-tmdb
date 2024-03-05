package tv

import (
	"encoding/json"
	"fmt"
)

func UnmarshalExternalIDsResp(data []byte) (ExternalIDsResp, error) {
	var r ExternalIDsResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type ExternalIDsResp struct {
	WikidataId  string
	FacebookId  string
	FreebaseId  string
	Id          int32
	ImdbId      string
	InstagramId string
	TvdbId      int32
	TwitterId   string
	TvrageId    int32
	FreebaseMid string
}

func (a Tv) ExternalIDs(series_id int32) (*ExternalIDsResp, error) {
	url := fmt.Sprintf("%s/3/tv/%d/external_ids", a.client.Cfg.Host, series_id)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalExternalIDsResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
