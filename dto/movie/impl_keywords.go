package movie

import (
	"encoding/json"
	"fmt"
)

type Keywordskeywords struct {
	Id   int32
	Name string
}

func UnmarshalKeywordsResp(data []byte) (KeywordsResp, error) {
	var r KeywordsResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type KeywordsResp struct {
	Id       int32
	Keywords []Keywordskeywords
}

func (a Movie) Keywords(movie_id string) (*KeywordsResp, error) {
	url := fmt.Sprintf("%s/3/movie/%s/keywords", a.client.Cfg.Host, movie_id)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalKeywordsResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
