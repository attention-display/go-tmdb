package movie

import (
	"encoding/json"
	"fmt"
)

type VideosReq struct {
	Language string
}

func (r VideosReq) Validate() error {
	return nil
}

func (r VideosReq) getUrlPath() string {
	res := ""
	if r.Language == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?language=%s", r.Language)
	}

	return res
}

type Videosresults struct {
	Official    bool
	Iso_3166_1  string
	Name        string
	PublishedAt string
	Type        string
	Id          string
	Site        string
	Size        int32
	Iso_639_1   string
	Key         string
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

func (a Movie) Videos(movie_id int32, req VideosReq) (*VideosResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/movie/%d/videos", a.client.Cfg.Host, movie_id)

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
