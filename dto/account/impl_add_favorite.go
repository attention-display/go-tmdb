package account

import (
	"fmt"

	"github.com/attention-display/go-tmdb/entity/errors"
	"github.com/attention-display/go-tmdb/entity/proto"
)

type AddFavoriteRequest struct {
	AccountId int32       `json:"account_id"`
	SessionId string      `json:"session_id"`
	RAW_BODY  interface{} `json:"RAW_BODY"`
}

func (d AddFavoriteRequest) getUrlPath() string {
	url := fmt.Sprintf("%d/%s", d.AccountId, "favorite")
	if d.SessionId != "" {
		return fmt.Sprintf("%s?session_id=%s", url, d.SessionId)
	}
	return url
}

func (a Account) AddFavorite(req AddFavoriteRequest) (*proto.CommonResponse, error) {
	url := fmt.Sprintf("%s/%s", a.getUrlPath(), req.getUrlPath())
	body, ok := req.RAW_BODY.([]byte)
	if !ok {
		return nil, errors.InvalidParamsErr("RAW_BODY is required.")
	}
	resp, err := a.client.Urlib.Post(url, a.client.Cfg.DefaultHeaders, body)
	if err != nil {
		return nil, err
	}
	res, err := proto.UnmarshalCommonResponse(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
