package qiita_test

import (
	"net/http"
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
	"github.com/stretchr/testify/assert"
	"github.com/yyh-gl/go-vcr-sample/qiita"
)

func Test_FetchUser(t *testing.T) {
	tests := []struct {
		testCase     string
		id           string
		wantLocation string
	}{
		{
			testCase:     "Qiitaからyyh-glのユーザ情報を取得できていること",
			id:           "yyh-gl",
			wantLocation: "Tokyo, Japan",
		},
	}

	// go-vcr のレコーダを生成
	// 通信内容は ../fixtures/qiita に保存される
	r, _ := recorder.New("../fixtures/qiita")
	defer r.Stop()

	customHTTPClient := &http.Client{
		Transport: r,
	}
	qiitaClient := qiita.NewClient(customHTTPClient)

	for _, tt := range tests {
		t.Run(tt.testCase, func(t *testing.T) {
			user := qiitaClient.FetchUser(tt.id)
			assert.Equal(t, tt.wantLocation, user.Location)
		})
	}
}
