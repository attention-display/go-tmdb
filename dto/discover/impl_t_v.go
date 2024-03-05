package discover

import (
	"encoding/json"
	"fmt"
)

type TVReq struct {
	Page                       int32
	WithType                   string
	WithOriginalLanguage       string
	WithStatus                 string
	WithoutCompanies           string
	WithoutWatchProviders      string
	VoteCountgte               int32
	VoteCountlte               int32
	WithKeywords               string
	WithoutKeywords            string
	WithWatchProviders         string
	WithoutGenres              string
	IncludeAdult               bool
	WatchRegion                string
	WithCompanies              string
	WithRuntimelte             int32
	FirstAirDateYear           int32
	Language                   string
	WithRuntimegte             int32
	WithNetworks               int32
	WithOriginCountry          string
	AirDategte                 string
	IncludeNullFirstAirDates   bool
	SortBy                     string
	WithGenres                 string
	AirDatelte                 string
	FirstAirDategte            string
	VoteAveragegte             int32
	VoteAveragelte             int32
	FirstAirDatelte            string
	ScreenedTheatrically       bool
	Timezone                   string
	WithWatchMonetizationTypes string
}

func (r TVReq) Validate() error {
	return nil
}

func (r TVReq) getUrlPath() string {
	res := ""
	if r.WithOriginCountry == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?with_origin_country=%s", r.WithOriginCountry)
	}

	if r.AirDategte == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?air_date.gte=%s", r.AirDategte)
	}

	if r.SortBy == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?sort_by=%s", r.SortBy)
	}

	if r.WithGenres == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?with_genres=%s", r.WithGenres)
	}

	if r.AirDatelte == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?air_date.lte=%s", r.AirDatelte)
	}

	if r.FirstAirDategte == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?first_air_date.gte=%s", r.FirstAirDategte)
	}

	if r.FirstAirDatelte == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?first_air_date.lte=%s", r.FirstAirDatelte)
	}

	if r.Timezone == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?timezone=%s", r.Timezone)
	}

	if r.WithWatchMonetizationTypes == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?with_watch_monetization_types=%s", r.WithWatchMonetizationTypes)
	}

	if r.WithType == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?with_type=%s", r.WithType)
	}

	if r.WithOriginalLanguage == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?with_original_language=%s", r.WithOriginalLanguage)
	}

	if r.WithStatus == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?with_status=%s", r.WithStatus)
	}

	if r.WithoutCompanies == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?without_companies=%s", r.WithoutCompanies)
	}

	if r.WithoutWatchProviders == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?without_watch_providers=%s", r.WithoutWatchProviders)
	}

	if r.WithKeywords == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?with_keywords=%s", r.WithKeywords)
	}

	if r.WithoutKeywords == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?without_keywords=%s", r.WithoutKeywords)
	}

	if r.WithWatchProviders == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?with_watch_providers=%s", r.WithWatchProviders)
	}

	if r.WithoutGenres == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?without_genres=%s", r.WithoutGenres)
	}

	if r.WatchRegion == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?watch_region=%s", r.WatchRegion)
	}

	if r.WithCompanies == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?with_companies=%s", r.WithCompanies)
	}

	if r.Language == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?language=%s", r.Language)
	}

	return res
}

type TVresults struct {
	OriginalLanguage string
	BackdropPath     string
	Popularity       int32
	PosterPath       string
	VoteCount        int32
	OriginCountry    []string
	OriginalName     string
	Name             string
	Overview         string
	FirstAirDate     string
	Id               int32
	VoteAverage      int32
	GenreIds         []int32
}

func UnmarshalTVResp(data []byte) (TVResp, error) {
	var r TVResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type TVResp struct {
	Results      []TVresults
	TotalPages   int32
	TotalResults int32
	Page         int32
}

func (a Discover) TV(req TVReq) (*TVResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/discover/tv", a.client.Cfg.Host)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalTVResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
