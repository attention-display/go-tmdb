package movie

import (
	"encoding/json"
	"fmt"
)

type AccountStatesReq struct {
	GuestSessionId string
	SessionId      string
}

func (r AccountStatesReq) Validate() error {
	return nil
}

func (r AccountStatesReq) getUrlPath() string {
	res := ""
	if r.GuestSessionId == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?guest_session_id=%s", r.GuestSessionId)
	}

	if r.SessionId == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?session_id=%s", r.SessionId)
	}

	return res
}

type AccountStatesrated struct {
	Value int32
}

func UnmarshalAccountStatesResp(data []byte) (AccountStatesResp, error) {
	var r AccountStatesResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type AccountStatesResp struct {
	Favorite  bool
	Id        int32
	Rated     AccountStatesrated
	Watchlist bool
}

func (a Movie) AccountStates(movie_id int32, req AccountStatesReq) (*AccountStatesResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/movie/%d/account_states", a.client.Cfg.Host, movie_id)

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
