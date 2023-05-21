package trierouter

import (
	"net/http"
	"testing"

	"github.com/lkeix/myrouter"
)

func defineRadixRouter() *myrouter.Router {
	r := myrouter.NewRouter()
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	r.GET("/", h)
	r.GET("/hoge", h)
	r.GET("/fuga", h)
	r.GET("/hoge/fuga", h)
	r.GET("/fuga/hoge", h)
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
