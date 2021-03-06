package github

import (
	"context"
	"fmt"
	"github.com/google/go-github/v33/github"
	"github.com/muhammednagy/pipedrive-challenge/config"
	"golang.org/x/oauth2"
	"net/http"
	"time"
)

// GetGists Will get all gists by a user in one requests if they are less than 100 or more than one request if they are more
// if you provide githubToken it will be able to do more requests before it gets rate limited
func GetGists(ctx context.Context, config config.Config, lastVisit *time.Time, username string) ([]*github.Gist, error) {
	tc := &http.Client{}
	if config.GithubToken != "" {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: config.GithubToken},
		)
		tc = oauth2.NewClient(ctx, ts)
	}
	client := github.NewClient(tc)

	gistListOptions := github.GistListOptions{}
	if lastVisit != nil {
		gistListOptions.Since = *lastVisit
	}

	gists, response, err := client.Gists.List(ctx, username, &gistListOptions)
	if response == nil || err != nil {
		return nil, fmt.Errorf("error occurred fetching user gists: %s", err)
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("error occurred fetching user gists: %s", response.Status)
	}

	if response.NextPage != 0 {
		pagesCount := response.LastPage
		for i := 2; i <= pagesCount; i++ {
			gistListOptions := github.GistListOptions{
				ListOptions: github.ListOptions{
					Page: i,
				},
			}
			if lastVisit != nil {
				gistListOptions.Since = *lastVisit
			}
			newGists, response, err := client.Gists.List(ctx, username, &gistListOptions)
			if response == nil || err != nil {
				return nil, fmt.Errorf("error occurred fetching user gists: %s", err)
			}
			if response.StatusCode != 200 {
				return nil, fmt.Errorf("error occurred fetching user gists: %s", response.Status)
			}
			gists = append(gists, newGists...)
		}
	}
	return gists, nil
}
