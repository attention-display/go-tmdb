package tv

import (
	"encoding/json"
	"fmt"
)

type Changesitems struct {
	Id     string
	Time   string
	Value  string
	Action string
}
type Changeschanges struct {
	Items []Changesitems
	Key   string
}

func UnmarshalChangesResp(data []byte) (ChangesResp, error) {
	var r ChangesResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type ChangesResp struct {
	Changes []Changeschanges
}

func (a Tv) Changes(episode_id int32) (*ChangesResp, error) {
	url := fmt.Sprintf("%s/3/tv/episode/%d/changes", a.client.Cfg.Host, episode_id)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalChangesResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
