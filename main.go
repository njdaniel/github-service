package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v42/github"
	"golang.org/x/oauth2"
)

func main() {
	token := os.Getenv("GITHUB_TOKEN")
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	repos, _, err := client.Repositories.List(ctx, "", nil)
	if err != nil {
		fmt.Printf("error listing github repositories: %v ", err)
	}
	for _, v := range repos {
		//fmt.Println(v.Name)
		fmt.Println(v.GetCloneURL())
	}
}
