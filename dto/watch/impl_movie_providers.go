package watch

import (
	"encoding/json"
	"fmt"
)

type MovieProvidersReq struct {
	Language    string
	WatchRegion string
}

func (r MovieProvidersReq) Validate() error {
	return nil
}

func (r MovieProvidersReq) getUrlPath() string {
	res := ""
	if r.WatchRegion == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?watch_region=%s", r.WatchRegion)
	}

	if r.Language == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?language=%s", r.Language)
	}

	return res
}

type MovieProvidersdisplayPriorities struct {
	SK int32
	BG int32
	BE int32
	AE int32
	PT int32
	IE int32
	LV int32
	GR int32
	DE int32
	GH int32
	PL int32
	HU int32
	EE int32
	LT int32
	AR int32
	RU int32
	AU int32
	TW int32
	MZ int32
	MX int32
	AT int32
	FI int32
	PE int32
	FR int32
	NZ int32
	NL int32
	ZA int32
	CH int32
	US int32
	CV int32
	BR int32
	PH int32
	SE int32
	SI int32
	GT int32
	JP int32
	MY int32
	UG int32
	CR int32
	CO int32
	EG int32
	IL int32
	NO int32
	ID int32
	TR int32
	DK int32
	ES int32
	MU int32
	PY int32
	IN int32
	BO int32
	CA int32
	CZ int32
	GB int32
	HK int32
	EC int32
	VE int32
	HN int32
	SG int32
	IT int32
	CL int32
	TH int32
	SA int32
}
type MovieProvidersresults struct {
	DisplayPriority   int32
	LogoPath          string
	ProviderId        int32
	ProviderName      string
	DisplayPriorities MovieProvidersdisplayPriorities
}

func UnmarshalMovieProvidersResp(data []byte) (MovieProvidersResp, error) {
	var r MovieProvidersResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type MovieProvidersResp struct {
	Results []MovieProvidersresults
}

func (a Watch) MovieProviders(req MovieProvidersReq) (*MovieProvidersResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/watch/providers/movie", a.client.Cfg.Host)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalMovieProvidersResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
