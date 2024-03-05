package watch

import (
	"encoding/json"
	"fmt"
)

type TVProvidersReq struct {
	Language    string
	WatchRegion string
}

func (r TVProvidersReq) Validate() error {
	return nil
}

func (r TVProvidersReq) getUrlPath() string {
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

type TVProvidersdisplayPriorities struct {
	ZA int32
	SE int32
	CL int32
	BG int32
	CA int32
	CV int32
	US int32
	AU int32
	FR int32
	EG int32
	IT int32
	DE int32
	NZ int32
	EC int32
	CO int32
	MY int32
	AT int32
	BR int32
	DK int32
	ID int32
	NO int32
	IL int32
	TH int32
	TR int32
	PT int32
	HN int32
	HU int32
	VE int32
	GR int32
	LT int32
	TW int32
	PH int32
	BO int32
	EE int32
	IN int32
	GT int32
	JP int32
	MZ int32
	IE int32
	PE int32
	CH int32
	PY int32
	AE int32
	AR int32
	GB int32
	CR int32
	HK int32
	NL int32
	FI int32
	ES int32
	GH int32
	LV int32
	RU int32
	CZ int32
	MX int32
	SK int32
	PL int32
	SI int32
	UG int32
	SG int32
	BE int32
	MU int32
	SA int32
}
type TVProvidersresults struct {
	ProviderName      string
	DisplayPriorities TVProvidersdisplayPriorities
	DisplayPriority   int32
	LogoPath          string
	ProviderId        int32
}

func UnmarshalTVProvidersResp(data []byte) (TVProvidersResp, error) {
	var r TVProvidersResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type TVProvidersResp struct {
	Results []TVProvidersresults
}

func (a Watch) TVProviders(req TVProvidersReq) (*TVProvidersResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/watch/providers/tv", a.client.Cfg.Host)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalTVProvidersResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
