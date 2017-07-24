package Middleware

import (
	"net/http"
	"github.com/NiciiA/AuthRest/application/service"
	"encoding/json"
	"github.com/NiciiA/AuthRest/application/domain"
	"github.com/NiciiA/AuthRest/application/jwt"
)

func AdminMiddleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ak := r.Header.Get("Authorization")
		t, err := Service.AuthorizationHeaderValidator(ak)

		if err != nil {
			w.WriteHeader(401)
			json.NewEncoder(w).Encode(Domain.JsonError{401, 6})
			return
		}

		l := JWT.IsLoggedIn(t)
		if !l {
			w.WriteHeader(401)
			json.NewEncoder(w).Encode(Domain.JsonError{401, 6})
			return
		}

		c := JWT.TokenClaims(t)
		if c["role"].(string) != "administrator" {
			w.WriteHeader(403)
			json.NewEncoder(w).Encode(Domain.JsonError{403, 7})
			return
		}

		next.ServeHTTP(w, r)
		return
	})
}