
package dto

import (
	"github.com/attention-display/go-tmdb/conf"
	"github.com/attention-display/go-tmdb/dto/trending"
	"github.com/attention-display/go-tmdb/dto/network"
	"github.com/attention-display/go-tmdb/dto/find"
	"github.com/attention-display/go-tmdb/dto/person"
	"github.com/attention-display/go-tmdb/dto/account"
	"github.com/attention-display/go-tmdb/dto/movie"
	"github.com/attention-display/go-tmdb/dto/guest_session"
	"github.com/attention-display/go-tmdb/dto/configuration"
	"github.com/attention-display/go-tmdb/dto/collection"
	"github.com/attention-display/go-tmdb/dto/review"
	"github.com/attention-display/go-tmdb/dto/authentication"
	"github.com/attention-display/go-tmdb/dto/genre"
	"github.com/attention-display/go-tmdb/dto/tv"
	"github.com/attention-display/go-tmdb/dto/watch"
	"github.com/attention-display/go-tmdb/dto/discover"
	"github.com/attention-display/go-tmdb/dto/keyword"
	"github.com/attention-display/go-tmdb/dto/certification"
	"github.com/attention-display/go-tmdb/dto/company"
	"github.com/attention-display/go-tmdb/dto/list"
	"github.com/attention-display/go-tmdb/dto/search"
	"github.com/attention-display/go-tmdb/dto/credit"
)

type Controller struct {
	trending.Trending
	network.Network
	find.Find
	person.Person
	account.Account
	movie.Movie
	guest_session.GuestSession
	configuration.Configuration
	collection.Collection
	review.Review
	authentication.Authentication
	genre.Genre
	tv.Tv
	watch.Watch
	discover.Discover
	keyword.Keyword
	certification.Certification
	company.Company
	list.List
	search.Search
	credit.Credit
}

func NewController(cfg *conf.Configuration) Controller {
	return Controller{
		Trending: trending.NewTrending(cfg),
		Network: network.NewNetwork(cfg),
		Find: find.NewFind(cfg),
		Person: person.NewPerson(cfg),
		Account: account.NewAccount(cfg),
		Movie: movie.NewMovie(cfg),
		GuestSession: guest_session.NewGuestSession(cfg),
		Configuration: configuration.NewConfiguration(cfg),
		Collection: collection.NewCollection(cfg),
		Review: review.NewReview(cfg),
		Authentication: authentication.NewAuthentication(cfg),
		Genre: genre.NewGenre(cfg),
		Tv: tv.NewTv(cfg),
		Watch: watch.NewWatch(cfg),
		Discover: discover.NewDiscover(cfg),
		Keyword: keyword.NewKeyword(cfg),
		Certification: certification.NewCertification(cfg),
		Company: company.NewCompany(cfg),
		List: list.NewList(cfg),
		Search: search.NewSearch(cfg),
		Credit: credit.NewCredit(cfg),
	}
}

