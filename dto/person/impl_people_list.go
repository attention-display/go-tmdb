package person

import (
	"encoding/json"
	"fmt"
)

type PeopleListReq struct {
	EndDate   string
	Page      int32
	StartDate string
}

func (r PeopleListReq) Validate() error {
	return nil
}

func (r PeopleListReq) getUrlPath() string {
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

type PeopleListresults struct {
	Adult bool
	Id    int32
}

func UnmarshalPeopleListResp(data []byte) (PeopleListResp, error) {
	var r PeopleListResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type PeopleListResp struct {
	TotalResults int32
	Page         int32
	Results      []PeopleListresults
	TotalPages   int32
}

func (a Person) PeopleList(req PeopleListReq) (*PeopleListResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/person/changes", a.client.Cfg.Host)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalPeopleListResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
