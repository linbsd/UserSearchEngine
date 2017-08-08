package github

import (
	"github.com/google/go-github/github"
	"context"
	"fmt"
)

func Users(user string)  {

	client := github.NewClient(nil)
	ctx := context.Background()

	opt := &github.SearchOptions{Sort:"url"}
	repos, _, err := client.Search.Users(ctx,"linbsd",opt)

	//if(nil != err){
		fmt.Print(repos,err)
	//}

}
