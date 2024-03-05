package account

import (
	"encoding/json"
	"fmt"
)

type FavoriteMoviesReq struct {
	Language  string
	Page      int32
	SessionId string
	SortBy    string
}

func (r FavoriteMoviesReq) Validate() error {
	return nil
}

func (r FavoriteMoviesReq) getUrlPath() string {
	res := ""
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

	if r.SessionId == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?session_id=%s", r.SessionId)
	}

	return res
}

type FavoriteMoviesresults struct {
	Popularity       int32
	PosterPath       string
	OriginalLanguage string
	Overview         string
	VoteAverage      int32
	VoteCount        int32
	Adult            bool
	BackdropPath     string
	ReleaseDate      string
	GenreIds         []int32
	OriginalTitle    string
	Title            string
	Video            bool
	Id               int32
}

func UnmarshalFavoriteMoviesResp(data []byte) (FavoriteMoviesResp, error) {
	var r FavoriteMoviesResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type FavoriteMoviesResp struct {
	Results      []FavoriteMoviesresults
	TotalPages   int32
	TotalResults int32
	Page         int32
}

func (a Account) FavoriteMovies(account_id int32, req FavoriteMoviesReq) (*FavoriteMoviesResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/account/%d/favorite/movies", a.client.Cfg.Host, account_id)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalFavoriteMoviesResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
