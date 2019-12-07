package main

import (
	"fmt"

	"github.com/yyh-gl/go-vcr-sample/qiita"
)

func main() {
	user := qiita.FetchUser("yyh-gl")
	fmt.Println("============ RESULT ============")
	fmt.Printf("%+v\n", user)
	fmt.Println("============ RESULT ============")
}
