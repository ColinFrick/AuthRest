package UserHandler

import (
	"net/http"
	"github.com/NiciiA/AuthRest/application/service"
	"github.com/NiciiA/AuthRest/application/jwt"
	"encoding/json"
	"github.com/NiciiA/AuthRest/application/domain"
	"github.com/NiciiA/AuthRest/application/dao"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

func List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")


	ak := r.Header.Get("Authorization")
	t, _ := Service.AuthorizationHeaderValidator(ak)
	c := JWT.TokenClaims(t)
	if c["role"].(string) != "administrator" {
		w.WriteHeader(403)
		json.NewEncoder(w).Encode(Domain.JsonError{403, 7})
		return
	}

	o := r.URL.Query().Get("offset")
	l := r.URL.Query().Get("limit")

	if o == "" {
		o = "0"
	}
	if l == "" {
		l = "20"
	}

	oN, erN := Service.NumberURIValidator(o)
	oL, erL := Service.NumberURIValidator(l)

	if erN != nil || erL != nil || !oN || !oL {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(Domain.JsonError{400, 8})
		return
	}

	cL := Domain.Users{}
	oR, _ := strconv.Atoi(o)
	lR, _ := strconv.Atoi(l)
	Dao.GetUsersCollection().Find(bson.M{}).Skip(oR).Limit(lR).All(cL)

	json.NewEncoder(w).Encode(cL)
	w.WriteHeader(200)
	return
}
