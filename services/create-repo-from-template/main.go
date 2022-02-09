package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v42/github"
	"golang.org/x/oauth2"
)

var (
	name        = flag.String("name", "", "Name of repo to create in authenticated user's GitHub account.")
	description = flag.String("description", "", "Description of created repo.")
	private     = flag.Bool("private", false, "Will created repo be private.")
)

func main() {
	flag.Parse()
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Fatal("Unauthorized: No token present")
	}
	if *name == "" {
		log.Fatal("No name: New repos must be given a name")
	}
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	//https://github.com/njdaniel/api-template
	//get template repo
	template, _, err := client.Repositories.Get(ctx, "njdaniel", "api-template")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T\n", template)

	//r := &github.Repository{Name: name, Private: private, Description: description, TemplateRepository: template}
	owner := "njdaniel"
	t := &github.TemplateRepoRequest{Name: name, Owner: &owner, Private: private}
	repo, _, err := client.Repositories.CreateFromTemplate(ctx, "njdaniel", "api-template", t)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Successfully created new repo: %v\n", repo.GetName())
}
