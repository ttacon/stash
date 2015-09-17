package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"

	"github.com/codegangsta/cli"
	"github.com/olekukonko/tablewriter"
	"github.com/ttacon/stash"
)

func main() {

	app := cli.NewApp()
	app.Name = "clash"
	app.Usage = "clash - a cli for stash"
	app.Action = func(c *cli.Context) {
		fmt.Println("please choose a subcommand")
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "user, u",
			Value: "",
			Usage: "user to authenticate to stash with",
		},
		cli.StringFlag{
			Name:  "password, p",
			Value: "",
			Usage: "password to authenticate to stash with",
		},
		cli.StringFlag{
			Name:  "host",
			Value: "",
			Usage: "host stash instance is running on",
		},
	}

	app.Commands = []cli.Command{
		groupCommand(),
		pullRequestCommand(),
	}

	app.Run(os.Args)
}

func isValidUser(c *stash.Client, username string) bool {
	user, err := c.UserService().GetUser(username)
	if err != nil {
		return false
	}
	return !user.IsEmpty()
}

func groupCommand() cli.Command {
	return cli.Command{
		Name:  "groups",
		Usage: "manipulate stash groups",
		Subcommands: []cli.Command{
			cli.Command{
				Name:  "list",
				Usage: "list users in group",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "group, g",
						Value: "",
						Usage: "group to list users in",
					},
				},
				Action: func(c *cli.Context) {
					client := stash.NewClient(
						c.GlobalString("u"),
						c.GlobalString("p"),
						c.GlobalString("host"),
					)
					groupName := c.String("g")
					users, err := client.GroupService().GetUsers(groupName, "")
					if err != nil {
						fmt.Println("failed to interact with stash server at")
						return
					} else if len(users) == 0 {
						fmt.Println("there are no users in the group", groupName)
						return
					}

					displayUsers(users)
				},
			},
			cli.Command{
				Name:  "all",
				Usage: "list all groups",
				Action: func(c *cli.Context) {
					client := stash.NewClient(
						c.GlobalString("u"),
						c.GlobalString("p"),
						c.GlobalString("host"),
					)
					groups, err := client.GroupService().GetGroups("")
					if err != nil {
						fmt.Println("failed to interact with stash server at")
						return
					} else if len(groups) == 0 {
						fmt.Println("there are no groups :(")
						return
					}

					displayGroups(groups)
				},
			},
		},
	}
}

func displayGroup(group *stash.Group) {
	fmt.Printf("%s (is deleteable: %v)\n", group.Name, group.Deleteable)
}

func displayUsers(users []*stash.User) {
	var buf = bytes.NewBuffer(nil)
	table := tablewriter.NewWriter(buf)
	table.SetHeader(userHeaders)
	for _, user := range users {
		table.Append([]string{strconv.Itoa(user.ID), user.Name, user.EmailAddress})
	}
	table.Render() // Send output

	fmt.Println(buf.String())
}

var userHeaders = []string{"ID", "Name", "Email"}

func displayGroups(groups []*stash.Group) {
	var buf = bytes.NewBuffer(nil)
	table := tablewriter.NewWriter(buf)
	table.SetHeader(groupHeaders)

	for _, group := range groups {
		table.Append([]string{group.Name, strconv.FormatBool(group.Deleteable)})
	}

	table.Render()
	fmt.Println(buf.String())
}

var groupHeaders = []string{"Name", "Deleteable"}
