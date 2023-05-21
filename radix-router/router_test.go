package radixrouter

import (
	"net/http"
	"testing"

	"github.com/lkeix/myrouter"
)

func defineRadixRouter() *myrouter.Router {
	r := myrouter.NewRouter()
	r.GET("/", func(ctx *myrouter.Context) {})
	r.GET("/hoge", func(ctx *myrouter.Context) {})
	r.GET("/fuga", func(ctx *myrouter.Context) {})
	r.GET("/hoge/fuga", func(ctx *myrouter.Context) {})
	r.GET("/fuga/hoge", func(ctx *myrouter.Context) {})
	return r
}

func BenchmarkRadixRouter(b *testing.B) {
	b.ReportAllocs()
	r := defineRadixRouter()
	b.ResetTimer()

	routes := []string{
		"/",
		"/hoge",
		"/fuga",
		"/hoge/fuga",
		"/fuga/hoge",
	}

	for i := 0; i < b.N; i++ {
		for i := 0; i < len(routes); i++ {
			r.Search(http.MethodGet, routes[i])
		}
	}
}
