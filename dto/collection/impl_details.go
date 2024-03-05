package collection

import (
	"encoding/json"
	"fmt"
)

type DetailsReq struct {
	Language string
}

func (r DetailsReq) Validate() error {
	return nil
}

func (r DetailsReq) getUrlPath() string {
	res := ""
	if r.Language == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?language=%s", r.Language)
	}

	return res
}

type Detailsparts struct {
	PosterPath       string
	MediaType        string
	ReleaseDate      string
	BackdropPath     string
	Adult            bool
	VoteAverage      int32
	Video            bool
	GenreIds         []int32
	OriginalTitle    string
	Overview         string
	Popularity       int32
	VoteCount        int32
	Id               int32
	Title            string
	OriginalLanguage string
}

func UnmarshalDetailsResp(data []byte) (DetailsResp, error) {
	var r DetailsResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type DetailsResp struct {
	BackdropPath string
	Id           int32
	Name         string
	Overview     string
	Parts        []Detailsparts
	PosterPath   string
}

func (a Collection) Details(collection_id int32, req DetailsReq) (*DetailsResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/collection/%d", a.client.Cfg.Host, collection_id)

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
