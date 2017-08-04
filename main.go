package main

import (
	"fmt"
	"github.com/larrylv/go-weibo/weibo"
	"user-search-engine/stackoverflow"
)
func main() {
	fmt.Print("Hello,World\n")
	var accessToken = ""

	accessToken = "access_token"
	client := weibo.NewClient(accessToken)
	print(client.BaseURL.Host)

	// list  weibo
	opts := &weibo.StatusListOptions{}
	status, _, err := client.Statuses.UserTimeline(opts)
	fmt.Printf("Status is :\n",status)
	fmt.Printf("Err is :\n",err)
	fmt.Print("\n")

	//test stackoverflow stackoverflow.ge)tAnswers(

	stackoverflow.GetUsers()
}
