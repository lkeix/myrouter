package myrouter

import (
	"net/http"
	"reflect"
	"testing"
)

func TestRouter_GET(t *testing.T) {
	testcases := []struct {
		name     string
		endpoint string
		handler  http.Handler
	}{
		{
			name:     "/のエンドポイントにハンドラを追加する",
			endpoint: "/",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			}),
		},
		{
			name:     "/helloのエンドポイントにハンドラを追加する",
			endpoint: "/hello",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			}),
		},
		{
			name:     "/hogeのエンドポイントにハンドラを追加する",
			endpoint: "/hoge",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			}),
		},
		{
			name:     "/hoge/fugaのエンドポイントにハンドラを追加する",
			endpoint: "/hoge/fuga",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			}),
		},
	}

	r := NewRouter()

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			// panicが起きないことを確認する
			r.GET(testcase.endpoint, testcase.handler)
		})
	}
}

func TestRouter_Search(t *testing.T) {
	testcases := []struct {
		name     string
		endpoint string
		handler  http.Handler
	}{
		{
			name:     "/のエンドポイントのハンドラを取得する",
			endpoint: "/",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			}),
		},
		{
			name:     "/hoge/fugaのエンドポイントのハンドラを取得する",
			endpoint: "/hoge/fuga",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			}),
		},
		{
			name:     "/helloのエンドポイントのハンドラを取得する",
			endpoint: "/hello",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			}),
		},
		{
			name:     "/hogeのエンドポイントのハンドラを取得する",
			endpoint: "/hoge",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			}),
		},
		{
			name:     "/piyoのエンドポイントのハンドラを取得する",
			endpoint: "/piyo",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			}),
		},
		{
			name:     "/piyo/hogeのエンドポイントのハンドラを取得する",
			endpoint: "/piyo/hoge",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			}),
		},
		{
			name:     "/piyo/fugaのエンドポイントのハンドラを取得する",
			endpoint: "/piyo/fuga",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			}),
		},
	}

	r := NewRouter()

	for _, testcase := range testcases {
		r.GET(testcase.endpoint, testcase.handler)
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			handler := r.Search(http.MethodGet, testcase.endpoint)
			// 関数のポインタを比較する
			if reflect.ValueOf(handler).Pointer() != reflect.ValueOf(testcase.handler).Pointer() {
				t.Errorf("ハンドラが異なります\nexpected: %v\nactual: %v", testcase.handler, handler)
			}
		})
	}
}
