package movie

import (
	"encoding/json"
	"fmt"
)

type DetailsReq struct {
	AppendToResponse string
	Language         string
}

func (r DetailsReq) Validate() error {
	return nil
}

func (r DetailsReq) getUrlPath() string {
	res := ""
	if r.AppendToResponse == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?append_to_response=%s", r.AppendToResponse)
	}

	if r.Language == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?language=%s", r.Language)
	}

	return res
}

type DetailsspokenLanguages struct {
	Name        string
	EnglishName string
	Iso_639_1   string
}
type DetailsproductionCountries struct {
	Iso_3166_1 string
	Name       string
}
type Detailsgenres struct {
	Id   int32
	Name string
}
type DetailsproductionCompanies struct {
	Name          string
	OriginCountry string
	Id            int32
	LogoPath      string
}

func UnmarshalDetailsResp(data []byte) (DetailsResp, error) {
	var r DetailsResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type DetailsResp struct {
	ImdbId              string
	Overview            string
	PosterPath          string
	Title               string
	BackdropPath        string
	Homepage            string
	Video               bool
	ProductionCountries []DetailsproductionCountries
	Revenue             int32
	ReleaseDate         string
	Status              string
	Budget              int32
	SpokenLanguages     []DetailsspokenLanguages
	Runtime             int32
	VoteCount           int32
	Popularity          int32
	ProductionCompanies []DetailsproductionCompanies
	Id                  int32
	OriginalLanguage    string
	Genres              []Detailsgenres
	Tagline             string
	VoteAverage         int32
	Adult               bool
	OriginalTitle       string
}

func (a Movie) Details(movie_id int32, req DetailsReq) (*DetailsResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/movie/%d", a.client.Cfg.Host, movie_id)

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
