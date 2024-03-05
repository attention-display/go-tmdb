package collection

import (
	"encoding/json"
	"fmt"
)

type Translationsdata struct {
	Homepage string
	Overview string
	Title    string
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

func (a Collection) Translations(collection_id int32) (*TranslationsResp, error) {
	url := fmt.Sprintf("%s/3/collection/%d/translations", a.client.Cfg.Host, collection_id)

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
