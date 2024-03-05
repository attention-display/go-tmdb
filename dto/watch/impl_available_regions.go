package watch

import (
	"encoding/json"
	"fmt"
)

type AvailableRegionsReq struct {
	Language string
}

func (r AvailableRegionsReq) Validate() error {
	return nil
}

func (r AvailableRegionsReq) getUrlPath() string {
	res := ""
	if r.Language == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?language=%s", r.Language)
	}

	return res
}

type AvailableRegionsresults struct {
	EnglishName string
	Iso_3166_1  string
	NativeName  string
}

func UnmarshalAvailableRegionsResp(data []byte) (AvailableRegionsResp, error) {
	var r AvailableRegionsResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type AvailableRegionsResp struct {
	Results []AvailableRegionsresults
}

func (a Watch) AvailableRegions(req AvailableRegionsReq) (*AvailableRegionsResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/watch/providers/regions", a.client.Cfg.Host)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalAvailableRegionsResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
