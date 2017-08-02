package stackoverflow


//address
var api string = "https://api.stackexchange.com"
//urls
var answers_url string = "/2.2/answers?order=desc&sort=activity&site=stackoverflow"

func GetApi() string {
	return api
}
func GetAnswersURL() string {
	return answers_url
}

