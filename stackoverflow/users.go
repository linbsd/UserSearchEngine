package stackoverflow

import (
	"os"
	"fmt"
	"github.com/gorilla/http"
)

func GetUsers()  {
	rs, err := http.Get(os.Stdout, api + version + users_url + "?" + "order=" + order + "&sort=" + sort_type["r"] + "&site=" + site)

	if err != nil {
		fmt.Print(rs)
	}
}
