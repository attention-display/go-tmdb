package tv

import (
	"encoding/json"
	"fmt"
)

type VideosReq struct {
	Language             string
	IncludeVideoLanguage string
}

func (r VideosReq) Validate() error {
	return nil
}

func (r VideosReq) getUrlPath() string {
	res := ""
	if r.IncludeVideoLanguage == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?include_video_language=%s", r.IncludeVideoLanguage)
	}

	if r.Language == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?language=%s", r.Language)
	}

	return res
}

type Videosresults struct {
	Size        int32
	Id          string
	Key         string
	Official    bool
	PublishedAt string
	Type        string
	Iso_639_1   string
	Name        string
	Iso_3166_1  string
	Site        string
}

func UnmarshalVideosResp(data []byte) (VideosResp, error) {
	var r VideosResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type VideosResp struct {
	Id      int32
	Results []Videosresults
}

func (a Tv) Videos(series_id int32, req VideosReq) (*VideosResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/tv/%d/videos", a.client.Cfg.Host, series_id)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalVideosResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
