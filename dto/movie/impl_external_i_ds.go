package movie

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
	FacebookId string
	Id         int32
	ImdbId     string
}

func (a Movie) ExternalIDs(movie_id int32) (*ExternalIDsResp, error) {
	url := fmt.Sprintf("%s/3/movie/%d/external_ids", a.client.Cfg.Host, movie_id)

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
