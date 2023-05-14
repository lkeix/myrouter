package myrouter

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestRouter_GET(t *testing.T) {
	testcases := []struct {
		name     string
		endpoint string
		handler  CustomHandler
	}{
		{
			name:     "/のエンドポイントにハンドラを追加する",
			endpoint: "/",
			handler: func(c *Context) {

			},
		},
		{
			name:     "/helloのエンドポイントにハンドラを追加する",
			endpoint: "/hello",
			handler: func(c *Context) {

			},
		},
		{
			name:     "/hogeのエンドポイントにハンドラを追加する",
			endpoint: "/hoge",
			handler: func(c *Context) {

			},
		},
		{
			name:     "/hoge/fugaのエンドポイントにハンドラを追加する",
			endpoint: "/hoge/fuga",
			handler: func(c *Context) {

			},
		},
		{
			name:     "/:userのエンドポイントにハンドラを追加する",
			endpoint: "/hoge/:user",
			handler: func(c *Context) {

			},
		},
		{
			name:     "/:userのエンドポイントにハンドラを追加する",
			endpoint: "/:user/hoge",
			handler: func(c *Context) {

			},
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
		name           string
		setendpoint    string
		accessendpoint string
		handler        CustomHandler
	}{
		{
			name:           "/のエンドポイントのハンドラを取得する",
			setendpoint:    "/",
			accessendpoint: "/",
			handler: func(c *Context) {

			},
		},
		{
			name:           "/hoge/fugaのエンドポイントのハンドラを取得する",
			setendpoint:    "/hoge/fuga",
			accessendpoint: "/hoge/fuga",
			handler: func(c *Context) {

			},
		},
		{
			name:           "/helloのエンドポイントのハンドラを取得する",
			setendpoint:    "/hello",
			accessendpoint: "/hello",
			handler: func(c *Context) {

			},
		},
		{
			name:           "/hogeのエンドポイントのハンドラを取得する",
			setendpoint:    "/hoge",
			accessendpoint: "/hoge",
			handler: func(c *Context) {

			},
		},
		{
			name:           "/piyoのエンドポイントのハンドラを取得する",
			setendpoint:    "/piyo",
			accessendpoint: "/piyo",
			handler: func(c *Context) {

			},
		},
		{
			name:           "/piyo/hogeのエンドポイントのハンドラを取得する",
			setendpoint:    "/piyo/hoge",
			accessendpoint: "/piyo/hoge",
			handler: func(c *Context) {

			},
		},
		{
			name:           "/piyo/fugaのエンドポイントのハンドラを取得する",
			setendpoint:    "/piyo/fuga",
			accessendpoint: "/piyo/fuga",
			handler: func(c *Context) {

			},
		},
		{
			name:           "/hoge/:profileのエンドポイントにハンドラを追加する",
			setendpoint:    "/hoge/:profile",
			accessendpoint: "/hoge/www",
			handler: func(c *Context) {

			},
		},
		{
			name:           "/:foo/hogeのエンドポイントにハンドラを追加する",
			setendpoint:    "/:foo/hoge",
			accessendpoint: "/www/hoge",
			handler: func(c *Context) {

			},
		},
		{
			name:           "/:foo/hoge/:barのエンドポイントにハンドラを追加する",
			setendpoint:    "/:foo/hoge/:bar",
			accessendpoint: "/fwww/hoge/bwww",
			handler: func(c *Context) {

			},
		},
		{
			name:           "/:foo/:profileのエンドポイントにハンドラを追加する",
			setendpoint:    "/:foo/:profile",
			accessendpoint: "/fwww/pwww",
			handler: func(c *Context) {

			},
		},

		{
			name:           "/:foo/:profile/:barのエンドポイントにハンドラを追加する",
			setendpoint:    "/:foo/:profile/:bar",
			accessendpoint: "/fwww/pwww/bwww",
			handler: func(c *Context) {

			},
		},
	}

	r := NewRouter()

	for _, testcase := range testcases {
		r.GET(testcase.setendpoint, testcase.handler)
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			handler, params := r.Search(http.MethodGet, testcase.accessendpoint)

			for _, p := range params {
				fmt.Printf("key: %s\nvalue: %s\n", p.key, p.value)
			}

			// 関数のポインタを比較する
			if reflect.ValueOf(handler).Pointer() != reflect.ValueOf(testcase.handler).Pointer() {
				t.Errorf("ハンドラが異なります\nexpected: %v\nactual: %v", testcase.handler, handler)
			}
		})
	}
}

func TestRouter_ServeHTTP(t *testing.T) {
	router := NewRouter()
	testcases := []struct {
		name           string
		setendpoint    string
		accessendpoint string
		handler        CustomHandler
	}{
		{
			name:           "/のエンドポイントのハンドラを取得する",
			setendpoint:    "/",
			accessendpoint: "/",
			handler: func(c *Context) {

			},
		},
		{
			name:           "/hogeのエンドポイントのハンドラを取得する",
			setendpoint:    "/hoge",
			accessendpoint: "/hoge",
			handler: func(c *Context) {

			},
		},
		{
			name:           "/hoge/:profileのエンドポイントにハンドラを追加する",
			setendpoint:    "/hoge/:profile",
			accessendpoint: "/hoge/www1",
			handler: func(c *Context) {
				p := c.Param("profile")
				fmt.Println(p)
			},
		},
		{
			name:           "/:foo/hogeのエンドポイントにハンドラを追加する",
			setendpoint:    "/:foo/hoge",
			accessendpoint: "/www/hoge",
			handler: func(c *Context) {
				p := c.Param("foo")
				fmt.Println(p)
			},
		},
	}
	for _, testcase := range testcases {
		router.GET(testcase.setendpoint, testcase.handler)
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodGet, testcase.accessendpoint, nil)
			router.ServeHTTP(nil, req)
		})
	}
}
