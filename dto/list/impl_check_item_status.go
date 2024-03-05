package list

import (
	"encoding/json"
	"fmt"
)

type CheckItemStatusReq struct {
	Language string
	MovieId  int32
}

func (r CheckItemStatusReq) Validate() error {
	return nil
}

func (r CheckItemStatusReq) getUrlPath() string {
	res := ""
	if r.Language == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?language=%s", r.Language)
	}

	return res
}

func UnmarshalCheckItemStatusResp(data []byte) (CheckItemStatusResp, error) {
	var r CheckItemStatusResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type CheckItemStatusResp struct {
	Id          int32
	ItemPresent bool
}

func (a List) CheckItemStatus(list_id int32, req CheckItemStatusReq) (*CheckItemStatusResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/list/%d/item_status", a.client.Cfg.Host, list_id)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalCheckItemStatusResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
