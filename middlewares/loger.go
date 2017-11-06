package middlewares

import (
	"net/http"

	"github.com/pressly/chi/middleware"
)

var (
	DefaultLogger = middleware.DefaultLogger
)

func Logger() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			DefaultLogger(next).ServeHTTP(w, r)
			return
		}
		return http.HandlerFunc(fn)
	}
}
