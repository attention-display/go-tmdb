package account

import (
	"encoding/json"
	"fmt"
)

type RatedTVEpisodesReq struct {
	Language  string
	Page      int32
	SessionId string
	SortBy    string
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

	if r.SessionId == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?session_id=%s", r.SessionId)
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
	Rating         int32
	Runtime        int32
	SeasonNumber   int32
	Overview       string
	EpisodeNumber  int32
	VoteCount      int32
	StillPath      string
	VoteAverage    int32
	ProductionCode string
	ShowId         int32
	Id             int32
	Name           string
}

func UnmarshalRatedTVEpisodesResp(data []byte) (RatedTVEpisodesResp, error) {
	var r RatedTVEpisodesResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type RatedTVEpisodesResp struct {
	TotalResults int32
	Page         int32
	Results      []RatedTVEpisodesresults
	TotalPages   int32
}

func (a Account) RatedTVEpisodes(account_id int32, req RatedTVEpisodesReq) (*RatedTVEpisodesResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/account/%d/rated/tv/episodes", a.client.Cfg.Host, account_id)

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
