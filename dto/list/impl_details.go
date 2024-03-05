package list

import (
	"encoding/json"
	"fmt"
)

type DetailsReq struct {
	Language string
	Page     int32
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

type Detailsitems struct {
	OriginalTitle    string
	BackdropPath     string
	MediaType        string
	Title            string
	Adult            bool
	GenreIds         []int32
	OriginalLanguage string
	VoteAverage      int32
	Id               int32
	VoteCount        int32
	PosterPath       string
	Popularity       int32
	Video            bool
	Overview         string
	ReleaseDate      string
}

func UnmarshalDetailsResp(data []byte) (DetailsResp, error) {
	var r DetailsResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type DetailsResp struct {
	Iso_639_1     string
	FavoriteCount int32
	Id            string
	Items         []Detailsitems
	PosterPath    string
	CreatedBy     string
	ItemCount     int32
	Name          string
	Description   string
}

func (a List) Details(list_id int32, req DetailsReq) (*DetailsResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/list/%d", a.client.Cfg.Host, list_id)

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
