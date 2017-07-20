package main

import (
	"github.com/google/go-github/github"
	"fmt"
	"context"
)
func main() {
	client := github.NewClient(nil)
	ctx := context.Background()

	// list public repositories for org "github"
	opt := &github.RepositoryListByOrgOptions{Type: "public"}
	repos, _, err := client.Repositories.ListByOrg(ctx, "github", opt)

	fmt.Print(repos,err)
}
