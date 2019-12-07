package main

import (
	"fmt"
	"net/http"

	"github.com/yyh-gl/go-vcr-sample/qiita"
)

func main() {
	qiitaClient := qiita.NewClient(http.DefaultClient)
	user := qiitaClient.FetchUser("yyh-gl")
	fmt.Println("============ RESULT ============")
	fmt.Printf("%+v\n", user)
	fmt.Println("============ RESULT ============")
}
