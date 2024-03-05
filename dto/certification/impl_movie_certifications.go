package certification

import (
	"encoding/json"
	"fmt"
)

type MovieCertificationsMO struct {
	Meaning       string
	Order         int32
	Certification string
}
type MovieCertificationsPH struct {
	Order         int32
	Certification string
	Meaning       string
}
type MovieCertificationsFI struct {
	Certification string
	Meaning       string
	Order         int32
}
type MovieCertificationsIN struct {
	Order         int32
	Certification string
	Meaning       string
}
type MovieCertificationsSE struct {
	Meaning       string
	Order         int32
	Certification string
}
type MovieCertificationsSG struct {
	Certification string
	Meaning       string
	Order         int32
}
type MovieCertificationsTR struct {
	Meaning       string
	Order         int32
	Certification string
}
type MovieCertificationsVI struct {
	Order         int32
	Certification string
	Meaning       string
}
type MovieCertificationsBG struct {
	Meaning       string
	Order         int32
	Certification string
}
type MovieCertificationsNO struct {
	Certification string
	Meaning       string
	Order         int32
}
type MovieCertificationsCA struct {
	Meaning       string
	Order         int32
	Certification string
}
type MovieCertificationsGB struct {
	Certification string
	Meaning       string
	Order         int32
}
type MovieCertificationsGR struct {
	Certification string
	Meaning       string
	Order         int32
}
type MovieCertificationsHK struct {
	Certification string
	Meaning       string
	Order         int32
}
type MovieCertificationsID struct {
	Certification string
	Meaning       string
	Order         int32
}
type MovieCertificationsNL struct {
	Order         int32
	Certification string
	Meaning       string
}
type MovieCertificationsAR struct {
	Meaning       string
	Order         int32
	Certification string
}
type MovieCertificationsBR struct {
	Certification string
	Meaning       string
	Order         int32
}
type MovieCertificationsNZ struct {
	Certification string
	Meaning       string
	Order         int32
}
type MovieCertificationsIT struct {
	Certification string
	Meaning       string
	Order         int32
}
type MovieCertificationsLT struct {
	Certification string
	Meaning       string
	Order         int32
}
type MovieCertificationsCAQC struct {
	Certification string
	Meaning       string
	Order         int32
}
type MovieCertificationsFR struct {
	Certification string
	Meaning       string
	Order         int32
}
type MovieCertificationsZA struct {
	Certification string
	Meaning       string
	Order         int32
}
type MovieCertificationsPR struct {
	Certification string
	Meaning       string
	Order         int32
}
type MovieCertificationsPT struct {
	Certification string
	Meaning       string
	Order         int32
}
type MovieCertificationsLU struct {
	Meaning       string
	Order         int32
	Certification string
}
type MovieCertificationsLV struct {
	Certification string
	Meaning       string
	Order         int32
}
type MovieCertificationsDK struct {
	Meaning       string
	Order         int32
	Certification string
}
type MovieCertificationsHU struct {
	Certification string
	Meaning       string
	Order         int32
}
type MovieCertificationsJP struct {
	Order         int32
	Certification string
	Meaning       string
}
type MovieCertificationsMY struct {
	Certification string
	Meaning       string
	Order         int32
}
type MovieCertificationsTW struct {
	Certification string
	Meaning       string
	Order         int32
}
type MovieCertificationsDE struct {
	Certification string
	Meaning       string
	Order         int32
}
type MovieCertificationsIL struct {
	Order         int32
	Certification string
	Meaning       string
}
type MovieCertificationsES struct {
	Meaning       string
	Order         int32
	Certification string
}
type MovieCertificationsIE struct {
	Order         int32
	Certification string
	Meaning       string
}
type MovieCertificationsKR struct {
	Order         int32
	Certification string
	Meaning       string
}
type MovieCertificationsMX struct {
	Certification string
	Meaning       string
	Order         int32
}
type MovieCertificationsRU struct {
	Order         int32
	Certification string
	Meaning       string
}
type MovieCertificationsSK struct {
	Meaning       string
	Order         int32
	Certification string
}
type MovieCertificationsAU struct {
	Order         int32
	Certification string
	Meaning       string
}
type MovieCertificationsCH struct {
	Certification string
	Meaning       string
	Order         int32
}
type MovieCertificationsTH struct {
	Meaning       string
	Order         int32
	Certification string
}
type MovieCertificationsUS struct {
	Certification string
	Meaning       string
	Order         int32
}
type MovieCertificationscertifications struct {
	HK   []MovieCertificationsHK
	IN   []MovieCertificationsIN
	NO   []MovieCertificationsNO
	GR   []MovieCertificationsGR
	SG   []MovieCertificationsSG
	IL   []MovieCertificationsIL
	HU   []MovieCertificationsHU
	ES   []MovieCertificationsES
	AU   []MovieCertificationsAU
	FI   []MovieCertificationsFI
	BR   []MovieCertificationsBR
	FR   []MovieCertificationsFR
	LU   []MovieCertificationsLU
	TW   []MovieCertificationsTW
	IE   []MovieCertificationsIE
	MX   []MovieCertificationsMX
	US   []MovieCertificationsUS
	TR   []MovieCertificationsTR
	BG   []MovieCertificationsBG
	ZA   []MovieCertificationsZA
	NZ   []MovieCertificationsNZ
	PT   []MovieCertificationsPT
	LV   []MovieCertificationsLV
	DE   []MovieCertificationsDE
	KR   []MovieCertificationsKR
	MO   []MovieCertificationsMO
	SE   []MovieCertificationsSE
	VI   []MovieCertificationsVI
	CAQC []MovieCertificationsCAQC
	DK   []MovieCertificationsDK
	JP   []MovieCertificationsJP
	RU   []MovieCertificationsRU
	ID   []MovieCertificationsID
	AR   []MovieCertificationsAR
	LT   []MovieCertificationsLT
	PH   []MovieCertificationsPH
	IT   []MovieCertificationsIT
	MY   []MovieCertificationsMY
	PR   []MovieCertificationsPR
	SK   []MovieCertificationsSK
	CH   []MovieCertificationsCH
	TH   []MovieCertificationsTH
	CA   []MovieCertificationsCA
	GB   []MovieCertificationsGB
	NL   []MovieCertificationsNL
}

func UnmarshalMovieCertificationsResp(data []byte) (MovieCertificationsResp, error) {
	var r MovieCertificationsResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type MovieCertificationsResp struct {
	Certifications MovieCertificationscertifications
}

func (a Certification) MovieCertifications() (*MovieCertificationsResp, error) {
	url := fmt.Sprintf("%s/3/certification/movie/list", a.client.Cfg.Host)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalMovieCertificationsResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
