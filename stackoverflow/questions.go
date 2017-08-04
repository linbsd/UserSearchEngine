package stackoverflow

import (
	"os"
	"fmt"
	"github.com/gorilla/http"
)

func GetQuestions()  {
	rs, err := http.Get(os.Stdout, api + version + question_url + "?" + "order=" + order + "&sort=" + sort_type["a"] + "&site=" + site)

	if err != nil {
		fmt.Print(rs)
	}
}