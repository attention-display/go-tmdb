package account

import (
	"encoding/json"
	"fmt"
)

type DetailsReq struct {
	SessionId string
}

func (d DetailsReq) getUrlPath() string {
	if d.SessionId == "" {
		return ""
	} else {
		return fmt.Sprintf("?session_id=%s", d.SessionId)
	}
}

func UnmarshalAccountdetailsreply(data []byte) (Accountdetailsreply, error) {
	var r Accountdetailsreply
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Accountdetailsreply) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Accountdetailsreply struct {
	IncludeAdult bool   `json:"include_adult"`
	ID           int64  `json:"id"`
	ISO639_1     string `json:"iso_639_1"`
	ISO3166_1    string `json:"iso_3166_1"`
	Name         string `json:"name"`
	Username     string `json:"username"`
	Avatar       Avatar `json:"avatar"`
}

type Avatar struct {
	Gravatar Gravatar `json:"gravatar"`
	Tmdb     Tmdb     `json:"tmdb"`
}

type Gravatar struct {
	Hash string `json:"hash"`
}

type Tmdb struct {
	AvatarPath string `json:"avatar_path"`
}

func (a Account) Details(accountId int32, opts DetailsReq) (*Accountdetailsreply, error) {
	url := fmt.Sprintf("%s/%d%s", a.getUrlPath(), accountId, opts.getUrlPath())
	resp, err := a.client.Urlib.Get(url, a.client.Cfg.DefaultHeaders)
	if err != nil {
		return nil, err
	}
	res, err := UnmarshalAccountdetailsreply(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
