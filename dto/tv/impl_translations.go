package tv

import (
	"encoding/json"
	"fmt"
)

type Translationsdata struct {
	Overview string
	Name     string
}
type Translationstranslations struct {
	Name        string
	Data        Translationsdata
	EnglishName string
	Iso_3166_1  string
	Iso_639_1   string
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

func (a Tv) Translations(series_id int32, season_number int32) (*TranslationsResp, error) {
	url := fmt.Sprintf("%s/3/tv/%d/season/%d/translations", a.client.Cfg.Host, series_id, season_number)

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
