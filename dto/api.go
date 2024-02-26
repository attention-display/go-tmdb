package dto

import (
	"github.com/attention-display/go-tmdb/dto/account"
	"github.com/attention-display/go-tmdb/dto/movies"
)

type Controller struct {
	movies.Movies
	account.Account
}
