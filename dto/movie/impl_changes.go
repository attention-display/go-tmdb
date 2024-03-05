package movie

import "fmt"

type ChangesReq struct {
	Page      int32
	StartDate string
	EndDate   string
}

func (r ChangesReq) Validate() error {
	return nil
}

func (r ChangesReq) getUrlPath() string {
	res := ""
	if r.EndDate == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?end_date=%s", r.EndDate)
	}

	if r.StartDate == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?start_date=%s", r.StartDate)
	}

	return res
}

func (a Movie) Changes(movie_id int32, req ChangesReq) (interface{}, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/movie/%d/changes", a.client.Cfg.Host, movie_id)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
