package keyword

import (
	"encoding/json"
	"fmt"
)

func UnmarshalDetailsResp(data []byte) (DetailsResp, error) {
	var r DetailsResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type DetailsResp struct {
	Name string
	Id   int32
}

func (a Keyword) Details(keyword_id int32) (*DetailsResp, error) {
	url := fmt.Sprintf("%s/3/keyword/%d", a.client.Cfg.Host, keyword_id)

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
