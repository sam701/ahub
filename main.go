package main

import (
	"os"

	"github.com/sam701/ahub/pr"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Usage = "Github goodies"
	app.Version = "0.1.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "team",
			Usage:  "Team name that following commands apply to",
			EnvVar: "AHUB_TEAM",
		},
		cli.StringFlag{
			Name:   "org",
			Usage:  "Organization name that following commands apply to",
			EnvVar: "AHUB_ORG",
		},
	}
	app.Commands = []cli.Command{
		cli.Command{
			Name:  "browse",
			Usage: "Open browser on different github pages",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "user, u",
					Usage: "User of the repo if different from the origin's user.",
				},
				cli.BoolFlag{
					Name:  "pulls, p",
					Usage: "Open pull requests page",
				},
				cli.BoolFlag{
					Name:  "issues, i",
					Usage: "Open issues page",
				},
			},
			Action: Browse,
		},
		cli.Command{
			Name:  "pr",
			Usage: "Manage pull requests",
			Subcommands: []cli.Command{
				cli.Command{
					Name:   "list",
					Usage:  "List pull requests",
					Action: pr.List,
					Flags: []cli.Flag{
						cli.BoolFlag{
							Name:  "all",
							Usage: "List all PRs. By default only open PRs are listed.",
						},
						cli.StringFlag{
							Name:  "source, s",
							Usage: "From what `REPOS` to list (team, watch)",
							Value: "team",
						},
					},
				},
			},
		},
	}
	app.Run(os.Args)
}
