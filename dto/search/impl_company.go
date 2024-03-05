package search

import (
	"encoding/json"
	"fmt"

	"github.com/attention-display/go-tmdb/entity/errors"
)

type CompanyReq struct {
	Page  int32
	Query string
}

func (r CompanyReq) Validate() error {
	if r.Query == "" {
		return errors.InvalidParamsErr("Query is required.")
	}
	return nil
}

func (r CompanyReq) getUrlPath() string {
	res := ""
	if r.Query == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?query=%s", r.Query)
	}

	return res
}

type Companyresults struct {
	OriginCountry string
	Id            int32
	LogoPath      string
	Name          string
}

func UnmarshalCompanyResp(data []byte) (CompanyResp, error) {
	var r CompanyResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type CompanyResp struct {
	Results      []Companyresults
	TotalPages   int32
	TotalResults int32
	Page         int32
}

func (a Search) Company(req CompanyReq) (*CompanyResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/search/company", a.client.Cfg.Host)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalCompanyResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
