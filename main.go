package main

import (
	"fmt"
	"github.com/larrylv/go-weibo/weibo"
)
func main() {
	fmt.Print("Hello,World\n")
	var accessToken = ""
	var opts  = ""

	accessToken = "access_token"
	client := weibo.NewClient(accessToken)

	// Update a weibo
	opts = &weibo.StatusRequest{Status: weibo.String("Hello, Weibo!")}
	status, _, err := client.Statuses.Create(opts)
}
