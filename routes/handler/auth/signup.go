package AuthHandler

import "net/http"

func HandleSignup(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	return
}