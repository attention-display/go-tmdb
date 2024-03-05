package search

import (
	"encoding/json"
	"fmt"

	"github.com/attention-display/go-tmdb/entity/errors"
)

type KeywordReq struct {
	Query string
	Page  int32
}

func (r KeywordReq) Validate() error {
	if r.Query == "" {
		return errors.InvalidParamsErr("Query is required.")
	}
	return nil
}

func (r KeywordReq) getUrlPath() string {
	res := ""
	if r.Query == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?query=%s", r.Query)
	}

	return res
}

type Keywordresults struct {
	Id   int32
	Name string
}

func UnmarshalKeywordResp(data []byte) (KeywordResp, error) {
	var r KeywordResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type KeywordResp struct {
	Page         int32
	Results      []Keywordresults
	TotalPages   int32
	TotalResults int32
}

func (a Search) Keyword(req KeywordReq) (*KeywordResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/search/keyword", a.client.Cfg.Host)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalKeywordResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
