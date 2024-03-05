package guest_session

import (
	"encoding/json"
	"fmt"
)

type RatedTVEpisodesReq struct {
	Language string
	Page     int32
	SortBy   string
}

func (r RatedTVEpisodesReq) Validate() error {
	return nil
}

func (r RatedTVEpisodesReq) getUrlPath() string {
	res := ""
	if r.Language == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?language=%s", r.Language)
	}

	if r.SortBy == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?sort_by=%s", r.SortBy)
	}

	return res
}

type RatedTVEpisodesresults struct {
	AirDate        string
	VoteAverage    int32
	VoteCount      int32
	StillPath      string
	ProductionCode string
	Rating         int32
	Runtime        int32
	SeasonNumber   int32
	Name           string
	Id             int32
	Overview       string
	EpisodeNumber  int32
	ShowId         int32
}

func UnmarshalRatedTVEpisodesResp(data []byte) (RatedTVEpisodesResp, error) {
	var r RatedTVEpisodesResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type RatedTVEpisodesResp struct {
	Page         int32
	Results      []RatedTVEpisodesresults
	TotalPages   int32
	TotalResults int32
}

func (a GuestSession) RatedTVEpisodes(guest_session_id string, req RatedTVEpisodesReq) (*RatedTVEpisodesResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/guest_session/%s/rated/tv/episodes", a.client.Cfg.Host, guest_session_id)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalRatedTVEpisodesResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
