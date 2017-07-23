package AuthHandler

import "net/http"

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	return
}