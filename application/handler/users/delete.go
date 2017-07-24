package UserHandler

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/NiciiA/AuthRest/application/service"
	"github.com/NiciiA/AuthRest/application/jwt"
	"github.com/NiciiA/AuthRest/application/domain"
	"github.com/NiciiA/AuthRest/application/dao"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	v := mux.Vars(r)
	id := v["id"]

	ak := r.Header.Get("Authorization")
	t, _ := Service.AuthorizationHeaderValidator(ak)
	c := JWT.TokenClaims(t)
	if c["role"].(string) == "administrator" || id == c["id"].(string) {
		Dao.GetUsersCollection().Remove(bson.M{"_id": bson.ObjectIdHex(id)})
		w.WriteHeader(204)
		return
	} else {
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(Domain.JsonError{401, 6})
		return
	}
}
