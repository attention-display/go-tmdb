package movie

import (
	"encoding/json"
	"fmt"
)

type NowPlayingReq struct {
	Page     int32
	Region   string
	Language string
}

func (r NowPlayingReq) Validate() error {
	return nil
}

func (r NowPlayingReq) getUrlPath() string {
	res := ""
	if r.Language == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?language=%s", r.Language)
	}

	if r.Region == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?region=%s", r.Region)
	}

	return res
}

type NowPlayingdates struct {
	Maximum string
	Minimum string
}
type NowPlayingresults struct {
	VoteAverage      int32
	PosterPath       string
	Id               int32
	ReleaseDate      string
	GenreIds         []int32
	OriginalLanguage string
	Title            string
	OriginalTitle    string
	Video            bool
	Adult            bool
	BackdropPath     string
	Overview         string
	Popularity       int32
	VoteCount        int32
}

func UnmarshalNowPlayingResp(data []byte) (NowPlayingResp, error) {
	var r NowPlayingResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type NowPlayingResp struct {
	TotalPages   int32
	TotalResults int32
	Dates        NowPlayingdates
	Page         int32
	Results      []NowPlayingresults
}

func (a Movie) NowPlaying(req NowPlayingReq) (*NowPlayingResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/movie/now_playing", a.client.Cfg.Host)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalNowPlayingResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
