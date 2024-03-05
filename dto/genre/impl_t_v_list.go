package genre

import (
	"encoding/json"
	"fmt"
)

type TVListReq struct {
	Language string
}

func (r TVListReq) Validate() error {
	return nil
}

func (r TVListReq) getUrlPath() string {
	res := ""
	if r.Language == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?language=%s", r.Language)
	}

	return res
}

type TVListgenres struct {
	Id   int32
	Name string
}

func UnmarshalTVListResp(data []byte) (TVListResp, error) {
	var r TVListResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type TVListResp struct {
	Genres []TVListgenres
}

func (a Genre) TVList(req TVListReq) (*TVListResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/genre/tv/list", a.client.Cfg.Host)

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
