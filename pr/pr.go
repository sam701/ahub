package pr

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/google/go-github/github"
	"github.com/sam701/ahub/client"
	"github.com/sam701/tcolor"
	"github.com/urfave/cli"
)

func List(ctx *cli.Context) error {
	source := ctx.String("source")
	if source == "team" {
		listPRsInTeamsRepos(ctx.GlobalString("team"), ctx.GlobalString("org"))
	} else if source == "watch" {
		listPRsInWatchedRepos(ctx.GlobalString("org"))
	} else {
		cli.ShowCommandHelp(ctx, "list")
	}
	return nil
}

func listPRsInTeamsRepos(teamName, org string) {
	c := client.New()
	teams, _, err := c.Organizations.ListUserTeams(context.Background(), nil)
	if err != nil {
		log.Fatalln("ERROR", err)
	}
	for _, team := range teams {
		if *team.Name == teamName {
			repos, _, err := c.Organizations.ListTeamRepos(context.Background(), *team.ID, &github.ListOptions{
				Page:    0,
				PerPage: 1000,
			})
			if err != nil {
				log.Fatalln("ERROR", err)
			}
			scanRepos(repos, org, c)
		}
	}
}

func listPRsInWatchedRepos(org string) {
	c := client.New()
	repos, _, err := c.Activity.ListWatched(context.Background(), "", &github.ListOptions{
		Page:    0,
		PerPage: 1000,
	})
	if err != nil {
		log.Fatalln("ERROR", err)
	}
	scanRepos(repos, org, c)
}

func scanRepos(repos []*github.Repository, org string, c *github.Client) {
	for _, repo := range repos {
		if strings.HasPrefix(*repo.FullName, org) {
			fmt.Printf("%-60s %s\n",
				tcolor.Colorize(*repo.FullName, tcolor.New().Bold().Foreground(tcolor.Green)),
				*repo.HTMLURL)
			prs, _, err := c.PullRequests.List(context.Background(), org, *repo.Name, nil)
			if err != nil {
				log.Fatalln("ERROR", err)
			}

			for _, pr := range prs {
				fmt.Printf("  %4d %s\n       %s\n",
					*pr.Number,
					tcolor.Colorize(*pr.Title, tcolor.New().Foreground(tcolor.BrightCyan)),
					*pr.HTMLURL)
			}
		}
	}
}
