package stackoverflow

import (
	"os"
	"fmt"
	"github.com/gorilla/http"
)

func GetQuestions()  {
	rs, err := http.Get(os.Stdout, api + question_url)

	if err != nil {
		fmt.Print(rs)
	}
}