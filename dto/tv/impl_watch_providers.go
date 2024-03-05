package tv

import (
	"encoding/json"
	"fmt"
)

type WatchProvidersReq struct {
	Language string
}

func (r WatchProvidersReq) Validate() error {
	return nil
}

func (r WatchProvidersReq) getUrlPath() string {
	res := ""
	if r.Language == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?language=%s", r.Language)
	}

	return res
}

type WatchProvidersflatrate struct {
	DisplayPriority int32
	LogoPath        string
	ProviderId      int32
	ProviderName    string
}
type WatchProvidersMY struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersVE struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersbuy struct {
	DisplayPriority int32
	LogoPath        string
	ProviderId      int32
	ProviderName    string
}
type WatchProvidersrent struct {
	LogoPath        string
	ProviderId      int32
	ProviderName    string
	DisplayPriority int32
}
type WatchProvidersJP struct {
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
	Link     string
	Rent     []WatchProvidersrent
}
type WatchProvidersMK struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersTZ struct {
	Link     string
	Flatrate []WatchProvidersflatrate
}
type WatchProvidersHR struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersJM struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersMZ struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersPL struct {
	Flatrate []WatchProvidersflatrate
	Link     string
	Rent     []WatchProvidersrent
}
type WatchProvidersAU struct {
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersBB struct {
	Link     string
	Flatrate []WatchProvidersflatrate
}
type WatchProvidersCI struct {
	Link     string
	Flatrate []WatchProvidersflatrate
}
type WatchProvidersHU struct {
	Link     string
	Flatrate []WatchProvidersflatrate
}
type WatchProvidersIQ struct {
	Link     string
	Flatrate []WatchProvidersflatrate
}
type WatchProvidersUY struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersAT struct {
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersCO struct {
	Link     string
	Flatrate []WatchProvidersflatrate
}
type WatchProvidersDE struct {
	Link     string
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
}
type WatchProvidersEG struct {
	Link     string
	Flatrate []WatchProvidersflatrate
}
type WatchProvidersGQ struct {
	Link     string
	Flatrate []WatchProvidersflatrate
}
type WatchProvidersID struct {
	Link     string
	Flatrate []WatchProvidersflatrate
}
type WatchProvidersIE struct {
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersCA struct {
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersDK struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersDO struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersSE struct {
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersGH struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersLB struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersNG struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersCL struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersCR struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersDZ struct {
	Link     string
	Flatrate []WatchProvidersflatrate
}
type WatchProvidersLY struct {
	Link     string
	Flatrate []WatchProvidersflatrate
}
type WatchProvidersSA struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersAR struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersBA struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersBE struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersSN struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersTT struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersfree struct {
	LogoPath        string
	ProviderId      int32
	ProviderName    string
	DisplayPriority int32
}
type WatchProvidersUS struct {
	Flatrate []WatchProvidersflatrate
	Free     []WatchProvidersfree
	Link     string
	Buy      []WatchProvidersbuy
}
type WatchProvidersKE struct {
	Link     string
	Flatrate []WatchProvidersflatrate
}
type WatchProvidersPE struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersPS struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersSC struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersCZ struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersES struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersIL struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersGB struct {
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersMD struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersNZ struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersRU struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersBR struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersBS struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersCH struct {
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersNE struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersPA struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersTR struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersZA struct {
	Link     string
	Flatrate []WatchProvidersflatrate
}
type WatchProvidersBG struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersFR struct {
	Link     string
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
}
type WatchProvidersMU struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersSG struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersSK struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersIT struct {
	Flatrate []WatchProvidersflatrate
	Link     string
	Buy      []WatchProvidersbuy
}
type WatchProvidersPH struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersPY struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersUG struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersPT struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersRS struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersSI struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersTH struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersBO struct {
	Link     string
	Flatrate []WatchProvidersflatrate
}
type WatchProvidersGT struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersSV struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersRO struct {
	Link     string
	Flatrate []WatchProvidersflatrate
}
type WatchProvidersTW struct {
	Link     string
	Flatrate []WatchProvidersflatrate
}
type WatchProvidersZM struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersKR struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersMX struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersNL struct {
	Link     string
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
}
type WatchProvidersGF struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersHK struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersHN struct {
	Link     string
	Flatrate []WatchProvidersflatrate
}
type WatchProvidersAE struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersEC struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersFI struct {
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersNO struct {
	Link     string
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
}
type WatchProvidersresults struct {
	BO WatchProvidersBO
	EG WatchProvidersEG
	GQ WatchProvidersGQ
	CA WatchProvidersCA
	DK WatchProvidersDK
	GH WatchProvidersGH
	PE WatchProvidersPE
	RU WatchProvidersRU
	SV WatchProvidersSV
	RO WatchProvidersRO
	GF WatchProvidersGF
	VE WatchProvidersVE
	JM WatchProvidersJM
	PL WatchProvidersPL
	GT WatchProvidersGT
	MK WatchProvidersMK
	CI WatchProvidersCI
	DE WatchProvidersDE
	TT WatchProvidersTT
	PY WatchProvidersPY
	HR WatchProvidersHR
	CO WatchProvidersCO
	SN WatchProvidersSN
	NE WatchProvidersNE
	CL WatchProvidersCL
	BA WatchProvidersBA
	ES WatchProvidersES
	ZA WatchProvidersZA
	MZ WatchProvidersMZ
	ID WatchProvidersID
	CR WatchProvidersCR
	LY WatchProvidersLY
	US WatchProvidersUS
	RS WatchProvidersRS
	HN WatchProvidersHN
	UG WatchProvidersUG
	MY WatchProvidersMY
	AU WatchProvidersAU
	AT WatchProvidersAT
	BG WatchProvidersBG
	FR WatchProvidersFR
	MU WatchProvidersMU
	IT WatchProvidersIT
	PT WatchProvidersPT
	FI WatchProvidersFI
	IE WatchProvidersIE
	SE WatchProvidersSE
	NG WatchProvidersNG
	BR WatchProvidersBR
	TW WatchProvidersTW
	DZ WatchProvidersDZ
	CH WatchProvidersCH
	PS WatchProvidersPS
	TH WatchProvidersTH
	HU WatchProvidersHU
	IQ WatchProvidersIQ
	BE WatchProvidersBE
	IL WatchProvidersIL
	PA WatchProvidersPA
	TR WatchProvidersTR
	SK WatchProvidersSK
	LB WatchProvidersLB
	GB WatchProvidersGB
	NZ WatchProvidersNZ
	BS WatchProvidersBS
	NO WatchProvidersNO
	TZ WatchProvidersTZ
	AR WatchProvidersAR
	KE WatchProvidersKE
	SC WatchProvidersSC
	CZ WatchProvidersCZ
	SG WatchProvidersSG
	HK WatchProvidersHK
	BB WatchProvidersBB
	DO WatchProvidersDO
	SA WatchProvidersSA
	ZM WatchProvidersZM
	KR WatchProvidersKR
	MX WatchProvidersMX
	EC WatchProvidersEC
	JP WatchProvidersJP
	UY WatchProvidersUY
	MD WatchProvidersMD
	PH WatchProvidersPH
	SI WatchProvidersSI
	NL WatchProvidersNL
	AE WatchProvidersAE
}

func UnmarshalWatchProvidersResp(data []byte) (WatchProvidersResp, error) {
	var r WatchProvidersResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type WatchProvidersResp struct {
	Id      int32
	Results WatchProvidersresults
}

func (a Tv) WatchProviders(series_id int32, season_number int32, req WatchProvidersReq) (*WatchProvidersResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/tv/%d/season/%d/watch/providers", a.client.Cfg.Host, series_id, season_number)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalWatchProvidersResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
