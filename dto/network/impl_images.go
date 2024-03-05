package network

import (
	"encoding/json"
	"fmt"
)

type Imageslogos struct {
	Id          string
	VoteAverage int32
	VoteCount   int32
	Width       int32
	AspectRatio int32
	FilePath    string
	FileType    string
	Height      int32
}

func UnmarshalImagesResp(data []byte) (ImagesResp, error) {
	var r ImagesResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type ImagesResp struct {
	Id    int32
	Logos []Imageslogos
}

func (a Network) Images(network_id int32) (*ImagesResp, error) {
	url := fmt.Sprintf("%s/3/network/%d/images", a.client.Cfg.Host, network_id)

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
