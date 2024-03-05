package movie

import (
	"encoding/json"
	"fmt"
)

type ReleaseDatesreleaseDates struct {
	Iso_639_1     string
	Note          string
	ReleaseDate   string
	Type          int32
	Certification string
	Descriptors   []interface{}
}
type ReleaseDatesresults struct {
	ReleaseDates []ReleaseDatesreleaseDates
	Iso_3166_1   string
}

func UnmarshalReleaseDatesResp(data []byte) (ReleaseDatesResp, error) {
	var r ReleaseDatesResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type ReleaseDatesResp struct {
	Id      int32
	Results []ReleaseDatesresults
}

func (a Movie) ReleaseDates(movie_id int32) (*ReleaseDatesResp, error) {
	url := fmt.Sprintf("%s/3/movie/%d/release_dates", a.client.Cfg.Host, movie_id)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalReleaseDatesResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
