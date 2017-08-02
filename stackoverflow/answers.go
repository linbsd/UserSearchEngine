package stackoverflow

import (
	"fmt"
	"github.com/gorilla/http"
	"os"
)


func GetAnswers()  {

	rs, err := http.Get(os.Stdout, api + answers_url)

	if err != nil {
		fmt.Print(rs)
	}

}
