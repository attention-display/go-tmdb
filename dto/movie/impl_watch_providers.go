package movie

import (
	"encoding/json"
	"fmt"
)

type WatchProvidersflatrate struct {
	DisplayPriority int32
	LogoPath        string
	ProviderId      int32
	ProviderName    string
}
type WatchProvidersGF struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersbuy struct {
	LogoPath        string
	ProviderId      int32
	ProviderName    string
	DisplayPriority int32
}
type WatchProvidersLB struct {
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersrent struct {
	ProviderId      int32
	ProviderName    string
	DisplayPriority int32
	LogoPath        string
}
type WatchProvidersAE struct {
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
	Link     string
	Rent     []WatchProvidersrent
}
type WatchProvidersBE struct {
	Link     string
	Rent     []WatchProvidersrent
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
}
type WatchProvidersBG struct {
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
	Link     string
	Rent     []WatchProvidersrent
}
type WatchProvidersDO struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersMY struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersSK struct {
	Flatrate []WatchProvidersflatrate
	Link     string
	Rent     []WatchProvidersrent
	Buy      []WatchProvidersbuy
}
type WatchProvidersAU struct {
	Link     string
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
}
type WatchProvidersBA struct {
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersIQ struct {
	Link     string
	Flatrate []WatchProvidersflatrate
}
type WatchProvidersMK struct {
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersMT struct {
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
	Link     string
	Rent     []WatchProvidersrent
}
type WatchProvidersOM struct {
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
	Link     string
	Rent     []WatchProvidersrent
}
type WatchProvidersYE struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersAR struct {
	Rent     []WatchProvidersrent
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersIL struct {
	Link     string
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
}
type WatchProvidersEE struct {
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
	Link     string
	Rent     []WatchProvidersrent
}
type WatchProvidersads struct {
	ProviderId      int32
	ProviderName    string
	DisplayPriority int32
	LogoPath        string
}
type WatchProvidersHR struct {
	Ads      []WatchProvidersads
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersTR struct {
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
	Link     string
	Rent     []WatchProvidersrent
}
type WatchProvidersCL struct {
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
	Link     string
	Rent     []WatchProvidersrent
}
type WatchProvidersDK struct {
	Flatrate []WatchProvidersflatrate
	Link     string
	Rent     []WatchProvidersrent
	Buy      []WatchProvidersbuy
}
type WatchProvidersKR struct {
	Link     string
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
}
type WatchProvidersPA struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersPE struct {
	Link     string
	Rent     []WatchProvidersrent
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
}
type WatchProvidersBB struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersDE struct {
	Flatrate []WatchProvidersflatrate
	Link     string
	Rent     []WatchProvidersrent
	Buy      []WatchProvidersbuy
}
type WatchProvidersCO struct {
	Rent     []WatchProvidersrent
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersES struct {
	Ads      []WatchProvidersads
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
	Link     string
	Rent     []WatchProvidersrent
}
type WatchProvidersFJ struct {
	Buy  []WatchProvidersbuy
	Link string
}
type WatchProvidersGB struct {
	Rent     []WatchProvidersrent
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersHN struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersID struct {
	Link     string
	Flatrate []WatchProvidersflatrate
}
type WatchProvidersBH struct {
	Link string
	Buy  []WatchProvidersbuy
}
type WatchProvidersBS struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersLV struct {
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersJP struct {
	Rent     []WatchProvidersrent
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersLI struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersRS struct {
	Flatrate []WatchProvidersflatrate
	Link     string
	Buy      []WatchProvidersbuy
}
type WatchProvidersFR struct {
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
	Link     string
	Rent     []WatchProvidersrent
}
type WatchProvidersIN struct {
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
	Link     string
	Rent     []WatchProvidersrent
}
type WatchProvidersNO struct {
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
	Link     string
	Rent     []WatchProvidersrent
}
type WatchProvidersPT struct {
	Rent     []WatchProvidersrent
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersSA struct {
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
	Link     string
	Rent     []WatchProvidersrent
}
type WatchProvidersSE struct {
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
	Link     string
	Rent     []WatchProvidersrent
}
type WatchProvidersCZ struct {
	Link     string
	Rent     []WatchProvidersrent
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
}
type WatchProvidersEG struct {
	Buy  []WatchProvidersbuy
	Link string
	Rent []WatchProvidersrent
}
type WatchProvidersGT struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersSG struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersTW struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersUG struct {
	Buy  []WatchProvidersbuy
	Link string
	Rent []WatchProvidersrent
}
type WatchProvidersBR struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersFI struct {
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
	Link     string
	Rent     []WatchProvidersrent
}
type WatchProvidersPL struct {
	Link     string
	Rent     []WatchProvidersrent
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
}
type WatchProvidersIE struct {
	Rent     []WatchProvidersrent
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersPK struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersRO struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersTH struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersGR struct {
	Flatrate []WatchProvidersflatrate
	Link     string
	Rent     []WatchProvidersrent
	Buy      []WatchProvidersbuy
}
type WatchProvidersMX struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersSV struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersTT struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersUS struct {
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
	Link     string
	Rent     []WatchProvidersrent
}
type WatchProvidersMU struct {
	Buy  []WatchProvidersbuy
	Link string
	Rent []WatchProvidersrent
}
type WatchProvidersPH struct {
	Link     string
	Flatrate []WatchProvidersflatrate
}
type WatchProvidersCV struct {
	Buy  []WatchProvidersbuy
	Link string
	Rent []WatchProvidersrent
}
type WatchProvidersIS struct {
	Flatrate []WatchProvidersflatrate
	Link     string
	Buy      []WatchProvidersbuy
}
type WatchProvidersMD struct {
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersNL struct {
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
	Link     string
	Rent     []WatchProvidersrent
}
type WatchProvidersNZ struct {
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersPY struct {
	Link     string
	Flatrate []WatchProvidersflatrate
}
type WatchProvidersBO struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersCA struct {
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
	Link     string
	Rent     []WatchProvidersrent
}
type WatchProvidersHK struct {
	Link     string
	Flatrate []WatchProvidersflatrate
}
type WatchProvidersKW struct {
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersLT struct {
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
	Link     string
	Rent     []WatchProvidersrent
}
type WatchProvidersMZ struct {
	Rent []WatchProvidersrent
	Buy  []WatchProvidersbuy
	Link string
}
type WatchProvidersQA struct {
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersSI struct {
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersAL struct {
	Buy  []WatchProvidersbuy
	Link string
}
type WatchProvidersCH struct {
	Link     string
	Rent     []WatchProvidersrent
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
}
type WatchProvidersVE struct {
	Link     string
	Rent     []WatchProvidersrent
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
}
type WatchProvidersJO struct {
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersPS struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersRU struct {
	Rent     []WatchProvidersrent
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersUY struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersZA struct {
	Flatrate []WatchProvidersflatrate
	Link     string
	Rent     []WatchProvidersrent
	Buy      []WatchProvidersbuy
}
type WatchProvidersEC struct {
	Link     string
	Rent     []WatchProvidersrent
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
}
type WatchProvidersIT struct {
	Flatrate []WatchProvidersflatrate
	Link     string
	Rent     []WatchProvidersrent
	Buy      []WatchProvidersbuy
}
type WatchProvidersGI struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersHU struct {
	Flatrate []WatchProvidersflatrate
	Link     string
	Rent     []WatchProvidersrent
	Buy      []WatchProvidersbuy
}
type WatchProvidersJM struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersSM struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersAT struct {
	Rent     []WatchProvidersrent
	Buy      []WatchProvidersbuy
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersCR struct {
	Flatrate []WatchProvidersflatrate
	Link     string
}
type WatchProvidersresults struct {
	SV WatchProvidersSV
	CH WatchProvidersCH
	CO WatchProvidersCO
	FR WatchProvidersFR
	NO WatchProvidersNO
	PK WatchProvidersPK
	BA WatchProvidersBA
	EE WatchProvidersEE
	CL WatchProvidersCL
	CR WatchProvidersCR
	PT WatchProvidersPT
	PL WatchProvidersPL
	TH WatchProvidersTH
	CV WatchProvidersCV
	OM WatchProvidersOM
	IL WatchProvidersIL
	PE WatchProvidersPE
	BH WatchProvidersBH
	AT WatchProvidersAT
	SI WatchProvidersSI
	AU WatchProvidersAU
	RS WatchProvidersRS
	SE WatchProvidersSE
	GT WatchProvidersGT
	MK WatchProvidersMK
	CZ WatchProvidersCZ
	BO WatchProvidersBO
	LT WatchProvidersLT
	IT WatchProvidersIT
	PA WatchProvidersPA
	LI WatchProvidersLI
	KW WatchProvidersKW
	VE WatchProvidersVE
	LV WatchProvidersLV
	SA WatchProvidersSA
	SG WatchProvidersSG
	IS WatchProvidersIS
	SK WatchProvidersSK
	ES WatchProvidersES
	GB WatchProvidersGB
	ID WatchProvidersID
	NL WatchProvidersNL
	JO WatchProvidersJO
	HU WatchProvidersHU
	BG WatchProvidersBG
	TW WatchProvidersTW
	MD WatchProvidersMD
	ZA WatchProvidersZA
	MU WatchProvidersMU
	MY WatchProvidersMY
	IQ WatchProvidersIQ
	DK WatchProvidersDK
	DE WatchProvidersDE
	KR WatchProvidersKR
	JP WatchProvidersJP
	HK WatchProvidersHK
	JM WatchProvidersJM
	PH WatchProvidersPH
	NZ WatchProvidersNZ
	MZ WatchProvidersMZ
	AL WatchProvidersAL
	MT WatchProvidersMT
	IN WatchProvidersIN
	EG WatchProvidersEG
	FI WatchProvidersFI
	PS WatchProvidersPS
	LB WatchProvidersLB
	BB WatchProvidersBB
	TT WatchProvidersTT
	UY WatchProvidersUY
	MX WatchProvidersMX
	PY WatchProvidersPY
	CA WatchProvidersCA
	RU WatchProvidersRU
	AE WatchProvidersAE
	YE WatchProvidersYE
	BS WatchProvidersBS
	IE WatchProvidersIE
	EC WatchProvidersEC
	RO WatchProvidersRO
	GF WatchProvidersGF
	DO WatchProvidersDO
	HR WatchProvidersHR
	TR WatchProvidersTR
	GI WatchProvidersGI
	HN WatchProvidersHN
	BR WatchProvidersBR
	GR WatchProvidersGR
	QA WatchProvidersQA
	US WatchProvidersUS
	SM WatchProvidersSM
	BE WatchProvidersBE
	AR WatchProvidersAR
	FJ WatchProvidersFJ
	UG WatchProvidersUG
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

func (a Movie) WatchProviders(movie_id int32) (*WatchProvidersResp, error) {
	url := fmt.Sprintf("%s/3/movie/%d/watch/providers", a.client.Cfg.Host, movie_id)

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
