package tv

import (
	"encoding/json"
	"fmt"
)

type TVListReq struct {
	StartDate string
	EndDate   string
	Page      int32
}

func (r TVListReq) Validate() error {
	return nil
}

func (r TVListReq) getUrlPath() string {
	res := ""
	if r.StartDate == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?start_date=%s", r.StartDate)
	}

	if r.EndDate == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?end_date=%s", r.EndDate)
	}

	return res
}

type TVListresults struct {
	Adult bool
	Id    int32
}

func UnmarshalTVListResp(data []byte) (TVListResp, error) {
	var r TVListResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type TVListResp struct {
	Results      []TVListresults
	TotalPages   int32
	TotalResults int32
	Page         int32
}

func (a Tv) TVList(req TVListReq) (*TVListResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/tv/changes", a.client.Cfg.Host)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalTVListResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
