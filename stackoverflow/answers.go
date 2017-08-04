package stackoverflow

import (
	"fmt"
	"github.com/gorilla/http"
	"os"
)


func GetAnswers()  {

	rs, err := http.Get(os.Stdout, api + version + answers_url +"?" + "order=" + order + "&sort=" + sort_type["a"] + "&site=" + site)

	if err != nil {
		fmt.Print(rs)
	}
}

