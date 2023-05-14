package myrouter

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var routes = []string{
	"/",
	"/hoge",
	"/fuga",
	"/health",
	"/hoge/fuga",
	"/n9te9",
	"/n9te9/hey",
	"/n9te9/hey/hello",
}

func BenchmarkMyRouter(b *testing.B) {
	b.ReportAllocs()
	r := httptest.NewRequest("GET", "/", nil)
	u := r.URL
	w := httptest.NewRecorder()
	myr := defineMyRoute()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < len(routes); i++ {
			r.Method = http.MethodGet
			u.Path = routes[i]
			myr.ServeHTTP(w, r)
		}
	}
}

type Myvalue struct {
	str     string
	handler CustomHandler
}

func defineMyRoute() *Router {
	values := []Myvalue{
		{"/", func(ctx *Context) {}},
		{"/hoge", func(ctx *Context) {}},
		{"/fuga", func(ctx *Context) {}},
		{"/hoge/fuga", func(ctx *Context) {}},
		{"/health", func(ctx *Context) {}},
		{"/hey", func(ctx *Context) {}},
		{"/:user", func(ctx *Context) {}},
		{"/:user/hey", func(ctx *Context) {}},
		{"/:user/hey/:group", func(ctx *Context) {}},
	}

	r := NewRouter()

	for _, v := range values {
		r.insert(http.MethodGet, v.str, v.handler)
	}
	return r
}
