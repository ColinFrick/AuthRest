package UserHandler

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/NiciiA/AuthRest/application/service"
	"github.com/NiciiA/AuthRest/application/jwt"
	"encoding/json"
	"github.com/NiciiA/AuthRest/application/domain"
	"github.com/NiciiA/AuthRest/application/dao"
	"gopkg.in/mgo.v2/bson"
)

func Single(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	v := mux.Vars(r)
	id := v["id"]

	ak := r.Header.Get("Authorization")
	t, _ := Service.AuthorizationHeaderValidator(ak)
	c := JWT.TokenClaims(t)
	if c["role"].(string) == "administrator" || id == c["id"].(string) {
		u := Domain.User{}
		Dao.GetUsersCollection().Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&u)
		json.NewEncoder(w).Encode(u)
		w.WriteHeader(200)
		return
	} else {
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(Domain.JsonError{401, 6})
		return
	}
}
