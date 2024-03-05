package person

import (
	"encoding/json"
	"fmt"
)

type ChangesReq struct {
	Page      int32
	StartDate string
	EndDate   string
}

func (r ChangesReq) Validate() error {
	return nil
}

func (r ChangesReq) getUrlPath() string {
	res := ""
	if r.EndDate == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?end_date=%s", r.EndDate)
	}

	if r.StartDate == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?start_date=%s", r.StartDate)
	}

	return res
}

type Changesitems struct {
	Time       string
	Value      string
	Action     string
	Id         string
	Iso_3166_1 string
	Iso_639_1  string
}
type Changeschanges struct {
	Key   string
	Items []Changesitems
}

func UnmarshalChangesResp(data []byte) (ChangesResp, error) {
	var r ChangesResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type ChangesResp struct {
	Changes []Changeschanges
}

func (a Person) Changes(person_id int32, req ChangesReq) (*ChangesResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/person/%d/changes", a.client.Cfg.Host, person_id)

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
