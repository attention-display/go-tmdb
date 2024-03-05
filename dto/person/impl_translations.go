package person

import (
	"encoding/json"
	"fmt"
)

type Translationsdata struct {
	Biography string
}
type Translationstranslations struct {
	Data        Translationsdata
	EnglishName string
	Iso_3166_1  string
	Iso_639_1   string
	Name        string
}

func UnmarshalTranslationsResp(data []byte) (TranslationsResp, error) {
	var r TranslationsResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type TranslationsResp struct {
	Id           int32
	Translations []Translationstranslations
}

func (a Person) Translations(person_id int32) (*TranslationsResp, error) {
	url := fmt.Sprintf("%s/3/person/%d/translations", a.client.Cfg.Host, person_id)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalTranslationsResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
