package movie

import "fmt"

type RecommendationsReq struct {
	Language string
	Page     int32
}

func (r RecommendationsReq) Validate() error {
	return nil
}

func (r RecommendationsReq) getUrlPath() string {
	res := ""
	if r.Language == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?language=%s", r.Language)
	}

	return res
}

func (a Movie) Recommendations(movie_id int32, req RecommendationsReq) (interface{}, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/movie/%d/recommendations", a.client.Cfg.Host, movie_id)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
