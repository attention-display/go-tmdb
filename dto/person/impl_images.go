package person

import "fmt"

func (a Person) Images(person_id int32) (interface{}, error) {
	url := fmt.Sprintf("%s/3/person/%d/images", a.client.Cfg.Host, person_id)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
