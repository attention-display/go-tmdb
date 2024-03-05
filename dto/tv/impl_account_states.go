package tv

import (
	"encoding/json"
	"fmt"
)

type AccountStatesReq struct {
	SessionId      string
	GuestSessionId string
}

func (r AccountStatesReq) Validate() error {
	return nil
}

func (r AccountStatesReq) getUrlPath() string {
	res := ""
	if r.SessionId == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?session_id=%s", r.SessionId)
	}

	if r.GuestSessionId == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?guest_session_id=%s", r.GuestSessionId)
	}

	return res
}

type AccountStatesrated struct {
	Value int32
}
type AccountStatesresults struct {
	EpisodeNumber int32
	Id            int32
	Rated         AccountStatesrated
}

func UnmarshalAccountStatesResp(data []byte) (AccountStatesResp, error) {
	var r AccountStatesResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type AccountStatesResp struct {
	Id      int32
	Results []AccountStatesresults
}

func (a Tv) AccountStates(series_id int32, season_number int32, req AccountStatesReq) (*AccountStatesResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/tv/%d/season/%d/account_states", a.client.Cfg.Host, series_id, season_number)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalAccountStatesResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
