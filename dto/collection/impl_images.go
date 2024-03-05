package collection

import (
	"encoding/json"
	"fmt"
)

type ImagesReq struct {
	IncludeImageLanguage string
	Language             string
}

func (r ImagesReq) Validate() error {
	return nil
}

func (r ImagesReq) getUrlPath() string {
	res := ""
	if r.IncludeImageLanguage == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?include_image_language=%s", r.IncludeImageLanguage)
	}

	if r.Language == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?language=%s", r.Language)
	}

	return res
}

type Imagesbackdrops struct {
	VoteCount   int32
	Width       int32
	AspectRatio int32
	FilePath    string
	Height      int32
	VoteAverage int32
}
type Imagesposters struct {
	VoteAverage int32
	VoteCount   int32
	Width       int32
	AspectRatio int32
	FilePath    string
	Height      int32
	Iso_639_1   string
}

func UnmarshalImagesResp(data []byte) (ImagesResp, error) {
	var r ImagesResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type ImagesResp struct {
	Backdrops []Imagesbackdrops
	Id        int32
	Posters   []Imagesposters
}

func (a Collection) Images(collection_id int32, req ImagesReq) (*ImagesResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/collection/%d/images", a.client.Cfg.Host, collection_id)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalImagesResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
