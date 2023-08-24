package proto

import "encoding/json"

func UnmarshalAccountdetailsreply(data []byte) (Accountdetailsreply, error) {
	var r Accountdetailsreply
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Accountdetailsreply) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Accountdetailsreply struct {
	Avatar       Avatar `json:"avatar"`
	ID           int64  `json:"id"`
	ISO639_1     string `json:"iso_639_1"`
	ISO3166_1    string `json:"iso_3166_1"`
	Name         string `json:"name"`
	IncludeAdult bool   `json:"include_adult"`
	Username     string `json:"username"`
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
