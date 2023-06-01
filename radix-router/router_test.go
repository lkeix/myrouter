package radixrouter

import (
	"net/http"
	"testing"

	"github.com/lkeix/myrouter"
)

func defineRadixRouter() *myrouter.Router {
	r := myrouter.NewRouter()
	r.GET("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	r.GET("/hoge", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	r.GET("/fuga", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	r.GET("/hoge/fuga", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	r.GET("/fuga/hoge", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
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
