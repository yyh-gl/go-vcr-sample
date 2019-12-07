package main

import (
	"fmt"
	"net/http"

	"github.com/dnaeon/go-vcr/recorder"
	"github.com/yyh-gl/go-vcr-sample/qiita"
)

func main() {
	// go-vcr のレコーダを生成
	// 通信内容は /fixtures/qiita に保存される
	r, _ := recorder.New("fixtures/qiita")
	defer r.Stop()

	originalHTTPClient := http.Client{
		Transport: r,
	}

	qiitaClient := qiita.NewClient(&originalHTTPClient)
	user := qiitaClient.FetchUser("yyh-gl")
	fmt.Println("============ RESULT ============")
	fmt.Printf("%+v\n", user)
	fmt.Println("============ RESULT ============")
}
