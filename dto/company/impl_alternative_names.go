package company

import (
	"encoding/json"
	"fmt"
)

type AlternativeNamesresults struct {
	Type string
	Name string
}

func UnmarshalAlternativeNamesResp(data []byte) (AlternativeNamesResp, error) {
	var r AlternativeNamesResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type AlternativeNamesResp struct {
	Id      int32
	Results []AlternativeNamesresults
}

func (a Company) AlternativeNames(company_id int32) (*AlternativeNamesResp, error) {
	url := fmt.Sprintf("%s/3/company/%d/alternative_names", a.client.Cfg.Host, company_id)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalAlternativeNamesResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
