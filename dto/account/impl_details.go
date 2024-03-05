package account

import (
	"encoding/json"
	"fmt"
)

type DetailsReq struct {
	SessionId string
}

func (r DetailsReq) Validate() error {
	return nil
}

func (r DetailsReq) getUrlPath() string {
	res := ""
	if r.SessionId == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?session_id=%s", r.SessionId)
	}

	return res
}

type Detailsgravatar struct {
	Hash string
}
type Detailstmdb struct {
	AvatarPath string
}
type Detailsavatar struct {
	Tmdb     Detailstmdb
	Gravatar Detailsgravatar
}

func UnmarshalDetailsResp(data []byte) (DetailsResp, error) {
	var r DetailsResp
	err := json.Unmarshal(data, &r)
	return r, err
}

type DetailsResp struct {
	Avatar       Detailsavatar
	Id           int32
	IncludeAdult bool
	Iso_3166_1   string
	Iso_639_1    string
	Name         string
	Username     string
}

func (a Account) Details(account_id int32, req DetailsReq) (*DetailsResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/account/%d", a.client.Cfg.Host, account_id)

	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalDetailsResp(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
