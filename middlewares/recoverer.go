package middlewares

import (
	"log"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/prometheus/client_golang/prometheus"

)

var (
	recoverCount = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "recover_count",
		Help: "recover count",
	})
	errorLog = log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
)

func init() {
	prometheus.MustRegister(recoverCount)
}

func Recoverer(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {
				recoverCount.Inc()
				errorLog.Println(rvr, string(debug.Stack()))
				//controls.Respond(w, r, nil, errs.CMServerFaild)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
