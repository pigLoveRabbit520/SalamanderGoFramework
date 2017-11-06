package controls

import (
	"net/http"
)

func ShowHome(w http.ResponseWriter, r *http.Request) {

	Respond(w, r, "Hello World", nil)
}