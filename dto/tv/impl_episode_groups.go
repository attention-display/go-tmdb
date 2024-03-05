package tv

import (
	"encoding/json"
	"fmt"
)

type EpisodeGroupsnetwork struct {
	Name          string
	OriginCountry string
	Id            int32
	LogoPath      string
}
type EpisodeGroupsresults struct {
	Type         int32
	Description  string
	EpisodeCount int32
	GroupCount   int32
	Id           string
	Name         string
	Network      EpisodeGroupsnetwork
}

func UnmarshalEpisodeGroupsResp(data []byte) (EpisodeGroupsResp, error) {
	var r EpisodeGroupsResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type EpisodeGroupsResp struct {
	Id      int32
	Results []EpisodeGroupsresults
}

func (a Tv) EpisodeGroups(series_id int32) (*EpisodeGroupsResp, error) {
	url := fmt.Sprintf("%s/3/tv/%d/episode_groups", a.client.Cfg.Host, series_id)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalEpisodeGroupsResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
