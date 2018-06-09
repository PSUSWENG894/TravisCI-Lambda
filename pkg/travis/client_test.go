package travis_test

import (
	"net/http"
	"net/http/httptest"
	"../travis"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
	client *travis.Client
	repoSlug string;
	apiToken string;
)

func setup() func() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	repoSlug = "test/slug"
	apiToken = "token"

	client, _ = travis.New(travis.BaseURL(server.URL),
		                   travis.RepoSlug(repoSlug),
		                   travis.ApiToken(apiToken))

	return func() {
		server.Close()
	}
}
