package account

import (
	"fmt"

	"github.com/attention-display/go-tmdb/entity/errors"
	"github.com/attention-display/go-tmdb/entity/proto"
)

type AddToWatchlistReq struct {
	SessionId string
}

func (r AddToWatchlistReq) Validate() error {
	return nil
}

func (r AddToWatchlistReq) getUrlPath() string {
	res := ""
	if r.SessionId == "" {
		res += ""
	} else {
		res += fmt.Sprintf("?session_id=%s", r.SessionId)
	}

	return res
}

type AddToWatchlistReqBody struct {
	RAW_BODY string
}

func (r AddToWatchlistReqBody) Validate() error {
	if r.RAW_BODY == "" {
		return errors.InvalidParamsErr("RAW_BODY is required.")
	}
	return nil
}

func (a Account) AddToWatchlist(account_id int32, req AddToWatchlistReq, reqBody AddToWatchlistReqBody) (*proto.CommonResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	if err := reqBody.Validate(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/3/account/%d/watchlist", a.client.Cfg.Host, account_id)

	resp, err := a.client.Urlib.Post(url, a.client.Cfg.DefaultHeaders, []byte(reqBody.RAW_BODY))
	if err != nil {
		return nil, err
	}
	res, err := proto.UnmarshalCommonResponse(resp)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
