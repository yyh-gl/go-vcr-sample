package qiita_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yyh-gl/go-vcr-sample/qiita"
)

func Test_FetchUser(t *testing.T) {
	tests := []struct {
		testCase string
		id string
		wantLocation string
	}{
		{
			testCase: "Qiitaからyyh-glのユーザ情報を取得できていること",
			id: "yyh-gl",
			wantLocation: "Tokyo, Japan",
		},
	}

	for _, tt := range tests {
		qiitaClient := qiita.NewClient(http.DefaultClient)

		t.Run(tt.testCase, func(t *testing.T) {
			user := qiitaClient.FetchUser(tt.id)
			assert.Equal(t, tt.wantLocation, user.Location)
		})
	}
}
