package main

import (
	"flag"
	"fmt"

	"github.com/ttacon/pretty"
	"github.com/ttacon/stash"
)

var (
	username = flag.String("u", "", "username")
	password = flag.String("p", "", "password")
	host     = flag.String("h", "http://localhost:7990", "host to query")

	user = flag.String("usr", "", "user to retrieve")
)

func main() {
	flag.Parse()
	c := stash.NewClient(*username, *password, *host)
	user, err := c.UserService().GetUser(*user)
	if err != nil {
		fmt.Println("err:", err)
		return
	} else if user.IsEmpty() {
		fmt.Println("no user was returned")
		return
	}

	pretty.Println(user)
}
