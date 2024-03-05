package search

import (
	"encoding/json"
	"fmt"

	"github.com/attention-display/go-tmdb/entity/errors"
)

type CollectionReq struct {
	Query        string
	IncludeAdult bool
	Language     string
	Page         int32
	Region       string
}

func (r CollectionReq) Validate() error {
	if r.Query == "" {
		return errors.InvalidParamsErr("Query is required.")
	}
	return nil
}

func (r CollectionReq) getUrlPath() string {
	res := ""
	if r.Query == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?query=%s", r.Query)
	}

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

type Collectionresults struct {
	BackdropPath     string
	Id               int32
	Name             string
	OriginalLanguage string
	OriginalName     string
	Overview         string
	PosterPath       string
	Adult            bool
}

func UnmarshalCollectionResp(data []byte) (CollectionResp, error) {
	var r CollectionResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type CollectionResp struct {
	TotalResults int32
	Page         int32
	Results      []Collectionresults
	TotalPages   int32
}

func (a Search) Collection(req CollectionReq) (*CollectionResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/search/collection", a.client.Cfg.Host)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalCollectionResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
