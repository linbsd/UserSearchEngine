package stackoverflow

import (
	"fmt"
	"github.com/gorilla/http"
	"os"
)

//address
var addr string = "https://api.stackexchange.com"
//url
var url string = "/2.2/answers?order=desc&sort=activity&site=stackoverflow"


func answers() {

	rs, err := http.Get(os.Stdout, addr + url)

	if err != nil {
		fmt.Print(rs)
	}

}
