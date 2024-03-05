package person

import (
	"encoding/json"
	"fmt"
)

type TaggedImagesReq struct {
	Page int32
}

func (r TaggedImagesReq) Validate() error {
	return nil
}

func (r TaggedImagesReq) getUrlPath() string {
	res := ""
	return res
}

type TaggedImagesmedia struct {
	Overview         string
	MediaType        string
	OriginalLanguage string
	OriginalTitle    string
	Popularity       int32
	Id               int32
	PosterPath       string
	GenreIds         []int32
	Title            string
	Video            bool
	Adult            bool
	BackdropPath     string
	ReleaseDate      string
	VoteAverage      int32
	VoteCount        int32
}
type TaggedImagesresults struct {
	Width       int32
	Iso_639_1   string
	Media       TaggedImagesmedia
	MediaType   string
	VoteCount   int32
	AspectRatio int32
	FilePath    string
	Height      int32
	ImageType   string
	Id          string
	VoteAverage int32
}

func UnmarshalTaggedImagesResp(data []byte) (TaggedImagesResp, error) {
	var r TaggedImagesResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type TaggedImagesResp struct {
	TotalPages   int32
	TotalResults int32
	Id           int32
	Page         int32
	Results      []TaggedImagesresults
}

func (a Person) TaggedImages(person_id int32, req TaggedImagesReq) (*TaggedImagesResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/person/%d/tagged_images", a.client.Cfg.Host, person_id)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalTaggedImagesResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
