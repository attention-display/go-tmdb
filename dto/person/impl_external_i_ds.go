package person

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
	FacebookId  string
	FreebaseMid string
	FreebaseId  string
	ImdbId      string
	InstagramId string
	Id          int32
	TiktokId    string
	TvrageId    int32
	TwitterId   string
	WikidataId  string
}

func (a Person) ExternalIDs(person_id int32) (*ExternalIDsResp, error) {
	url := fmt.Sprintf("%s/3/person/%d/external_ids", a.client.Cfg.Host, person_id)

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
