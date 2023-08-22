package gotmdb

import "net/http"

type TMDb struct {
	// TMDb api key
	apiKey string
	// TMDb access token
	accessToken string
	// http client
	http.Client
}

func (t *TMDb) get() {

}

func (t *TMDb) post() {

}

func (t *TMDb) delete() {

}
