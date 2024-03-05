package find

import (
	"encoding/json"
	"fmt"

	"github.com/attention-display/go-tmdb/entity/errors"
)

type FindByIDReq struct {
	Language       string
	ExternalSource string
}

func (r FindByIDReq) Validate() error {
	if r.ExternalSource == "" {
		return errors.InvalidParamsErr("ExternalSource is required.")
	}
	return nil
}

func (r FindByIDReq) getUrlPath() string {
	res := ""
	if r.ExternalSource == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?external_source=%s", r.ExternalSource)
	}

	if r.Language == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?language=%s", r.Language)
	}

	return res
}

type FindByIDmovieResults struct {
	Overview         string
	OriginalLanguage string
	BackdropPath     string
	ReleaseDate      string
	Title            string
	VoteCount        int32
	VoteAverage      int32
	Id               int32
	MediaType        string
	Adult            bool
	GenreIds         []int32
	Popularity       int32
	OriginalTitle    string
	PosterPath       string
	Video            bool
}

func UnmarshalFindByIDResp(data []byte) (FindByIDResp, error) {
	var r FindByIDResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type FindByIDResp struct {
	PersonResults    []interface{}
	TvEpisodeResults []interface{}
	TvResults        []interface{}
	TvSeasonResults  []interface{}
	MovieResults     []FindByIDmovieResults
}

func (a Find) FindByID(external_id string, req FindByIDReq) (*FindByIDResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/find/%s", a.client.Cfg.Host, external_id)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalFindByIDResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
