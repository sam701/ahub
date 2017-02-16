package client

import (
	"fmt"
	"net/http"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func New() *github.Client {
	token := os.Getenv("GITHUB_TOKEN")

	var httpClient *http.Client
	if token == "" {
		fmt.Println("WARNING: using unauthenticated client! Set GITHUB_TOKEN environment variable.")
	} else {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		httpClient = oauth2.NewClient(oauth2.NoContext, ts)
	}
	return github.NewClient(httpClient)
}
