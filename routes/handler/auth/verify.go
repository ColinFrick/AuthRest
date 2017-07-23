package AuthHandler

import "net/http"

func HandleVerify(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	return
}