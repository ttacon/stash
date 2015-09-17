package main

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/ttacon/pretty"
	"github.com/ttacon/stash"
)

func pullRequestCommand() cli.Command {
	return cli.Command{
		Name:  "prs",
		Usage: "manipulate pull-requests",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "repo",
				Value: "",
				Usage: "repo to make pull request for",
			},
			cli.StringFlag{
				Name:  "proj",
				Value: "",
				Usage: "project repo is based in",
			},
		},
		Subcommands: []cli.Command{
			cli.Command{
				Name:  "list",
				Usage: "list all prs for a given repo",
				Action: func(c *cli.Context) {
					fmt.Println("not done")
				},
			},
			cli.Command{
				Name:  "create",
				Usage: "create a pull request",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "branch",
						Value: "",
						Usage: "branch to create PR from",
					},
					cli.StringFlag{
						Name:  "title",
						Value: "",
						Usage: "title of pull request",
					},
					cli.StringFlag{
						Name:  "desc",
						Value: "",
						Usage: "description of PR",
					},
					cli.StringSliceFlag{
						Name:  "reviewers",
						Value: nil,
						Usage: "reviewers for PR",
					},
				},
				Action: func(c *cli.Context) {
					client := stash.NewClient(
						c.GlobalString("u"),
						c.GlobalString("p"),
						c.GlobalString("host"),
					)

					pr, err := client.PullRequestService().CreatePRFromSrcRef(
						c.GlobalString("proj"),
						c.GlobalString("repo"),
						c.String("branch"),
						c.String("title"),
						c.String("desc"),
						c.StringSlice("reviewers"),
					)
					if err != nil {
						fmt.Println("failed to create PR:", err)
						return
					}

					pretty.Println(pr)
				},
			},
		},
	}
}
