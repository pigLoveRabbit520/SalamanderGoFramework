package controls

import (
	"net/http"
	"github.com/go-chi/render"
)

func Respond(w http.ResponseWriter, r *http.Request, v interface{}, err error) {
	render.Respond(w, r, v)
}

func Redirect(w http.ResponseWriter, r *http.Request, status int, p string, v interface{}, err error) {
	w.Header().Set("Location", p)
	w.WriteHeader(status)
	w.Write([]byte("跳转中"))
}
