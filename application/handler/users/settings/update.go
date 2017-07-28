package UserSettingsHandler

import (
	"net/http"
	"github.com/NiciiA/AuthRest/application/domain"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/NiciiA/AuthRest/application/service"
	"github.com/NiciiA/AuthRest/application/jwt"
	"github.com/NiciiA/AuthRest/application/dao"
	"gopkg.in/mgo.v2/bson"
)

func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	b := Domain.Settings{}
	json.NewDecoder(r.Body).Decode(&b)

	rl, erl := Service.LanguageValidator(b.Language)
	if !rl || erl != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(Domain.JsonError{400, 4})
		return
	}

	v := mux.Vars(r)
	id := v["id"]

	ak := r.Header.Get("Authorization")
	t, _ := Service.AuthorizationHeaderValidator(ak)
	c := JWT.TokenClaims(t)
	if c["role"].(string) == "administrator" || id == c["id"].(string) {
		b.User = bson.ObjectIdHex(id)
		Dao.GetSettingsCollection().Update(bson.M{"_id": bson.ObjectIdHex(id)}, b)
		json.NewEncoder(w).Encode(b)
		w.WriteHeader(204)
		return
	} else {
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(Domain.JsonError{401, 6})
		return
	}
}