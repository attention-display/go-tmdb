package configuration

import (
	"encoding/json"
	"fmt"
)

func UnmarshalJobsResp(data []byte) (JobsResp, error) {
	var r JobsResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type JobsResp struct {
	Department string
	Jobs       []string
}

func (a Configuration) Jobs() (*JobsResp, error) {
	url := fmt.Sprintf("%s/3/configuration/jobs", a.client.Cfg.Host)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalJobsResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
