package tv

import (
	"encoding/json"
	"fmt"
)

type ListsReq struct {
	Language string
	Page     int32
}

func (r ListsReq) Validate() error {
	return nil
}

func (r ListsReq) getUrlPath() string {
	res := ""
	if r.Language == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?language=%s", r.Language)
	}

	return res
}

type Listsresults struct {
	FavoriteCount int32
	Id            int32
	Iso_3166_1    string
	Iso_639_1     string
	ItemCount     int32
	Name          string
	Description   string
}

func UnmarshalListsResp(data []byte) (ListsResp, error) {
	var r ListsResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type ListsResp struct {
	Id           int32
	Page         int32
	Results      []Listsresults
	TotalPages   int32
	TotalResults int32
}

func (a Tv) Lists(series_id int32, req ListsReq) (*ListsResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/tv/%d/lists", a.client.Cfg.Host, series_id)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalListsResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
