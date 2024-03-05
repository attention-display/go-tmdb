package trending

import (
	"encoding/json"
	"fmt"
)

type AllReq struct {
	Language string
}

func (r AllReq) Validate() error {
	return nil
}

func (r AllReq) getUrlPath() string {
	res := ""
	if r.Language == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?language=%s", r.Language)
	}

	return res
}

type Allresults struct {
	OriginalTitle    string
	ReleaseDate      string
	BackdropPath     string
	Overview         string
	PosterPath       string
	VoteCount        int32
	VoteAverage      int32
	Popularity       int32
	Title            string
	Adult            bool
	Id               int32
	OriginalLanguage string
	GenreIds         []int32
	Video            bool
	MediaType        string
}

func UnmarshalAllResp(data []byte) (AllResp, error) {
	var r AllResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type AllResp struct {
	Page         int32
	Results      []Allresults
	TotalPages   int32
	TotalResults int32
}

func (a Trending) All(time_window string, req AllReq) (*AllResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/trending/all/%s", a.client.Cfg.Host, time_window)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalAllResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
