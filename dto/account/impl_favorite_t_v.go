package account

import (
	"encoding/json"
	"fmt"
)

type FavoriteTVReq struct {
	SortBy    string
	Language  string
	Page      int32
	SessionId string
}

func (r FavoriteTVReq) Validate() error {
	return nil
}

func (r FavoriteTVReq) getUrlPath() string {
	res := ""
	if r.SessionId == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?session_id=%s", r.SessionId)
	}

	if r.SortBy == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?sort_by=%s", r.SortBy)
	}

	if r.Language == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?language=%s", r.Language)
	}

	return res
}

type FavoriteTVresults struct {
	OriginalName     string
	Overview         string
	Name             string
	VoteCount        int32
	FirstAirDate     string
	BackdropPath     string
	OriginalLanguage string
	VoteAverage      int32
	Adult            bool
	GenreIds         []int32
	OriginCountry    []string
	Id               int32
	Popularity       int32
	PosterPath       string
}

func UnmarshalFavoriteTVResp(data []byte) (FavoriteTVResp, error) {
	var r FavoriteTVResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type FavoriteTVResp struct {
	Page         int32
	Results      []FavoriteTVresults
	TotalPages   int32
	TotalResults int32
}

func (a Account) FavoriteTV(account_id int32, req FavoriteTVReq) (*FavoriteTVResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/account/%d/favorite/tv", a.client.Cfg.Host, account_id)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalFavoriteTVResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
