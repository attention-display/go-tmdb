package person

import (
	"encoding/json"
	"fmt"
)

type DetailsReq struct {
	AppendToResponse string
	Language         string
}

func (r DetailsReq) Validate() error {
	return nil
}

func (r DetailsReq) getUrlPath() string {
	res := ""
	if r.Language == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?language=%s", r.Language)
	}

	if r.AppendToResponse == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?append_to_response=%s", r.AppendToResponse)
	}

	return res
}

func UnmarshalDetailsResp(data []byte) (DetailsResp, error) {
	var r DetailsResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type DetailsResp struct {
	Biography          string
	PlaceOfBirth       string
	Id                 int32
	ProfilePath        string
	Gender             int32
	Name               string
	ImdbId             string
	AlsoKnownAs        []string
	KnownForDepartment string
	Popularity         int32
	Adult              bool
	Birthday           string
}

func (a Person) Details(person_id int32, req DetailsReq) (*DetailsResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/person/%d", a.client.Cfg.Host, person_id)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalDetailsResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
