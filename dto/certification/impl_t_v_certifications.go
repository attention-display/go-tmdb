package certification

import (
	"encoding/json"
	"fmt"
)

type TVCertificationsGB struct {
	Certification string
	Meaning       string
	Order         int32
}
type TVCertificationsGR struct {
	Meaning       string
	Order         int32
	Certification string
}
type TVCertificationsMX struct {
	Certification string
	Meaning       string
	Order         int32
}
type TVCertificationsPL struct {
	Order         int32
	Certification string
	Meaning       string
}
type TVCertificationsRU struct {
	Certification string
	Meaning       string
	Order         int32
}
type TVCertificationsTW struct {
	Certification string
	Meaning       string
	Order         int32
}
type TVCertificationsFI struct {
	Certification string
	Meaning       string
	Order         int32
}
type TVCertificationsIN struct {
	Meaning       string
	Order         int32
	Certification string
}
type TVCertificationsMA struct {
	Meaning       string
	Order         int32
	Certification string
}
type TVCertificationsMY struct {
	Order         int32
	Certification string
	Meaning       string
}
type TVCertificationsPH struct {
	Meaning       string
	Order         int32
	Certification string
}
type TVCertificationsUS struct {
	Meaning       string
	Order         int32
	Certification string
}
type TVCertificationsKR struct {
	Order         int32
	Certification string
	Meaning       string
}
type TVCertificationsTR struct {
	Certification string
	Meaning       string
	Order         int32
}
type TVCertificationsSE struct {
	Certification string
	Meaning       string
	Order         int32
}
type TVCertificationsTH struct {
	Order         int32
	Certification string
	Meaning       string
}
type TVCertificationsVI struct {
	Order         int32
	Certification string
	Meaning       string
}
type TVCertificationsAU struct {
	Certification string
	Meaning       string
	Order         int32
}
type TVCertificationsBR struct {
	Certification string
	Meaning       string
	Order         int32
}
type TVCertificationsDK struct {
	Order         int32
	Certification string
	Meaning       string
}
type TVCertificationsES struct {
	Certification string
	Meaning       string
	Order         int32
}
type TVCertificationsNO struct {
	Meaning       string
	Order         int32
	Certification string
}
type TVCertificationsPR struct {
	Meaning       string
	Order         int32
	Certification string
}
type TVCertificationsSG struct {
	Order         int32
	Certification string
	Meaning       string
}
type TVCertificationsAR struct {
	Certification string
	Meaning       string
	Order         int32
}
type TVCertificationsCA struct {
	Meaning       string
	Order         int32
	Certification string
}
type TVCertificationsCAQC struct {
	Order         int32
	Certification string
	Meaning       string
}
type TVCertificationsDE struct {
	Certification string
	Meaning       string
	Order         int32
}
type TVCertificationsID struct {
	Order         int32
	Certification string
	Meaning       string
}
type TVCertificationsBG struct {
	Certification string
	Meaning       string
	Order         int32
}
type TVCertificationsIL struct {
	Certification string
	Meaning       string
	Order         int32
}
type TVCertificationsIT struct {
	Certification string
	Meaning       string
	Order         int32
}
type TVCertificationsNL struct {
	Certification string
	Meaning       string
	Order         int32
}
type TVCertificationsPT struct {
	Certification string
	Meaning       string
	Order         int32
}
type TVCertificationsFR struct {
	Certification string
	Meaning       string
	Order         int32
}
type TVCertificationsHU struct {
	Order         int32
	Certification string
	Meaning       string
}
type TVCertificationsZA struct {
	Meaning       string
	Order         int32
	Certification string
}
type TVCertificationsLT struct {
	Certification string
	Meaning       string
	Order         int32
}
type TVCertificationsNZ struct {
	Certification string
	Meaning       string
	Order         int32
}
type TVCertificationsSK struct {
	Certification string
	Meaning       string
	Order         int32
}
type TVCertificationscertifications struct {
	PL   []TVCertificationsPL
	IN   []TVCertificationsIN
	VI   []TVCertificationsVI
	DK   []TVCertificationsDK
	CAQC []TVCertificationsCAQC
	PT   []TVCertificationsPT
	LT   []TVCertificationsLT
	KR   []TVCertificationsKR
	FR   []TVCertificationsFR
	MX   []TVCertificationsMX
	BR   []TVCertificationsBR
	SG   []TVCertificationsSG
	AR   []TVCertificationsAR
	IL   []TVCertificationsIL
	GR   []TVCertificationsGR
	TW   []TVCertificationsTW
	MA   []TVCertificationsMA
	TR   []TVCertificationsTR
	TH   []TVCertificationsTH
	ES   []TVCertificationsES
	NO   []TVCertificationsNO
	PR   []TVCertificationsPR
	BG   []TVCertificationsBG
	IT   []TVCertificationsIT
	NL   []TVCertificationsNL
	HU   []TVCertificationsHU
	RU   []TVCertificationsRU
	ID   []TVCertificationsID
	FI   []TVCertificationsFI
	PH   []TVCertificationsPH
	US   []TVCertificationsUS
	AU   []TVCertificationsAU
	ZA   []TVCertificationsZA
	MY   []TVCertificationsMY
	SE   []TVCertificationsSE
	SK   []TVCertificationsSK
	GB   []TVCertificationsGB
	CA   []TVCertificationsCA
	DE   []TVCertificationsDE
	NZ   []TVCertificationsNZ
}

func UnmarshalTVCertificationsResp(data []byte) (TVCertificationsResp, error) {
	var r TVCertificationsResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type TVCertificationsResp struct {
	Certifications TVCertificationscertifications
}

func (a Certification) TVCertifications() (*TVCertificationsResp, error) {
	url := fmt.Sprintf("%s/3/certification/tv/list", a.client.Cfg.Host)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalTVCertificationsResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
