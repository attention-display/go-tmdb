package discover

import (
	"encoding/json"
	"fmt"
)

type MovieReq struct {
	VoteCountgte               int32
	WithRuntimegte             int32
	Certificationgte           string
	Certificationlte           string
	CertificationCountry       string
	IncludeVideo               bool
	ReleaseDategte             string
	VoteAveragelte             int32
	WithReleaseType            int32
	Year                       int32
	Region                     string
	SortBy                     string
	VoteCountlte               int32
	WatchRegion                string
	WithGenres                 string
	WithPeople                 string
	IncludeAdult               bool
	WithCompanies              string
	WithCrew                   string
	WithOriginalLanguage       string
	WithRuntimelte             int32
	WithoutWatchProviders      string
	Page                       int32
	WithOriginCountry          string
	WithWatchProviders         string
	Certification              string
	PrimaryReleaseYear         int32
	PrimaryReleaseDatelte      string
	ReleaseDatelte             string
	WithCast                   string
	WithKeywords               string
	WithoutCompanies           string
	WithoutKeywords            string
	Language                   string
	WithWatchMonetizationTypes string
	WithoutGenres              string
	PrimaryReleaseDategte      string
	VoteAveragegte             int32
}

func (r MovieReq) Validate() error {
	return nil
}

func (r MovieReq) getUrlPath() string {
	res := ""
	if r.PrimaryReleaseDategte == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?primary_release_date.gte=%s", r.PrimaryReleaseDategte)
	}

	if r.Certificationgte == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?certification.gte=%s", r.Certificationgte)
	}

	if r.Certificationlte == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?certification.lte=%s", r.Certificationlte)
	}

	if r.CertificationCountry == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?certification_country=%s", r.CertificationCountry)
	}

	if r.ReleaseDategte == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?release_date.gte=%s", r.ReleaseDategte)
	}

	if r.Region == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?region=%s", r.Region)
	}

	if r.SortBy == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?sort_by=%s", r.SortBy)
	}

	if r.WatchRegion == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?watch_region=%s", r.WatchRegion)
	}

	if r.WithGenres == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?with_genres=%s", r.WithGenres)
	}

	if r.WithPeople == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?with_people=%s", r.WithPeople)
	}

	if r.WithCompanies == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?with_companies=%s", r.WithCompanies)
	}

	if r.WithCrew == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?with_crew=%s", r.WithCrew)
	}

	if r.WithOriginalLanguage == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?with_original_language=%s", r.WithOriginalLanguage)
	}

	if r.WithoutWatchProviders == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?without_watch_providers=%s", r.WithoutWatchProviders)
	}

	if r.WithWatchProviders == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?with_watch_providers=%s", r.WithWatchProviders)
	}

	if r.Certification == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?certification=%s", r.Certification)
	}

	if r.PrimaryReleaseDatelte == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?primary_release_date.lte=%s", r.PrimaryReleaseDatelte)
	}

	if r.ReleaseDatelte == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?release_date.lte=%s", r.ReleaseDatelte)
	}

	if r.WithCast == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?with_cast=%s", r.WithCast)
	}

	if r.WithKeywords == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?with_keywords=%s", r.WithKeywords)
	}

	if r.WithOriginCountry == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?with_origin_country=%s", r.WithOriginCountry)
	}

	if r.WithoutCompanies == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?without_companies=%s", r.WithoutCompanies)
	}

	if r.WithoutKeywords == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?without_keywords=%s", r.WithoutKeywords)
	}

	if r.Language == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?language=%s", r.Language)
	}

	if r.WithWatchMonetizationTypes == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?with_watch_monetization_types=%s", r.WithWatchMonetizationTypes)
	}

	if r.WithoutGenres == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?without_genres=%s", r.WithoutGenres)
	}

	return res
}

type Movieresults struct {
	Overview         string
	Id               int32
	ReleaseDate      string
	OriginalLanguage string
	Popularity       int32
	Video            bool
	VoteCount        int32
	OriginalTitle    string
	PosterPath       string
	Title            string
	VoteAverage      int32
	GenreIds         []int32
	BackdropPath     string
	Adult            bool
}

func UnmarshalMovieResp(data []byte) (MovieResp, error) {
	var r MovieResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type MovieResp struct {
	TotalPages   int32
	TotalResults int32
	Page         int32
	Results      []Movieresults
}

func (a Discover) Movie(req MovieReq) (*MovieResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/discover/movie", a.client.Cfg.Host)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalMovieResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
