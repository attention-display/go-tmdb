package dto

import (
	"github.com/attention-display/go-tmdb/conf"
	"github.com/attention-display/go-tmdb/dto/account"
	"github.com/attention-display/go-tmdb/dto/movies"
)

type Controller struct {
	movies.Movies
	account.Account
}

func NewController(cfg *conf.Configuration) Controller {
	return Controller{
		Movies:  movies.NewMovies(cfg),
		Account: account.NewAccount(cfg),
	}
}
