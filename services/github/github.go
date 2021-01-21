package github

import (
	"context"
	"github.com/google/go-github/v33/github"
	"github.com/labstack/gommon/log"
	"github.com/muhammednagy/pipedirve-challenge/db"
	"github.com/muhammednagy/pipedirve-challenge/models"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
	"net/http"
)

// Get Gists
// Will get all gists by a user in one requests if they are less than 100 or more than one request if they are more
// if you provide githubToken it will be able to do more requests before it gets rate limited
func GetGists(config models.Config, dbConnection *gorm.DB, username, githubToken string) []*github.Gist {
	var person models.Person
	tc := &http.Client{}
	ctx := context.Background()
	if config.GithubToken != "" {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: githubToken},
		)
		tc = oauth2.NewClient(ctx, ts)
	}
	client := github.NewClient(tc)

	people := db.GetPeople(dbConnection, username)
	if len(people) == 0 {
		_ = db.SavePerson(dbConnection, username)
	} else {
		person = people[0]
	}

	gists, response, err := client.Gists.List(ctx, username, &github.GistListOptions{
		Since: person.LastVisit,
	})
	if response == nil || err != nil {
		log.Errorf("error occurred fetching user gists", err)
		return nil
	}
	if response.StatusCode != 200 {
		log.Errorf("error occurred fetching user gists")
		return nil
	}

	if response.NextPage != 0 {
		pagesCount := response.LastPage - 1
		for i := 2; i < pagesCount; i++ {
			newGists, response, err := client.Gists.List(ctx, username, &github.GistListOptions{
				Since: person.LastVisit,
				ListOptions: github.ListOptions{
					Page: i,
				},
			})
			if response == nil || err != nil {
				log.Errorf("error occurred fetching user gists", err)
				return nil
			}
			if response.StatusCode != 200 {
				log.Errorf("error occurred fetching user gists")
				return nil
			}
			gists = append(gists, newGists...)
		}
	}
	return gists
}
