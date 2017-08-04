package stackoverflow


//address
var api string = "https://api.stackexchange.com"
//order type
var order string = "desc"
//sort type
var sort_type  = map[string] string {
	"a":"activity",
	"c":"creation",
	"v":"votes",
	"h":"hot",
	"w":"week",
	"m":"month",
	"r":"reputation",
	"n":"name",
	"mo":"modified",
	"p":"popular",
	"rl":"relevance",
}
//site
var site string = "stackoverflow"
//api version
var version string = "/2.2/"
//urls
var answers_url string = "answers"
var question_url string = "questions"
var users_url string = "users"

