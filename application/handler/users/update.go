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

func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	b := Domain.User{}
	json.NewDecoder(r.Body).Decode(&b)

	reg, err := Service.PhoneNumberValidator(b.PhoneNumber)
	if !reg || err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(Domain.JsonError{400, 1})
		return
	}

	rf, erf := Service.NameValidator(b.FirstName)
	if !rf || erf != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(Domain.JsonError{400, 4})
		return
	}

	rl, erl := Service.NameValidator(b.LastName)
	if !rl || erl != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(Domain.JsonError{400, 4})
		return
	}

	ro, ero := Service.RoleValidator(b.Role)
	if !ro || ero != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(Domain.JsonError{400, 9})
		return
	}

	v := mux.Vars(r)
	id := v["id"]

	ak := r.Header.Get("Authorization")
	t, _ := Service.AuthorizationHeaderValidator(ak)
	c := JWT.TokenClaims(t)
	if c["role"].(string) == "administrator" || id == c["id"].(string) {
		if c["role"].(string) == "customer" && b.Role == "administrator" {
			u := Domain.User{}
			Dao.GetUsersCollection().Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&u)
			if u.Disabled || u.Role != "administrator" {
				w.WriteHeader(403)
				json.NewEncoder(w).Encode(Domain.JsonError{403, 7})
				return
			}
		}
		b.ID = bson.ObjectIdHex(id)
		Dao.GetUsersCollection().Update(bson.M{"_id": bson.ObjectIdHex(id)}, b)
		json.NewEncoder(w).Encode(b)
		w.WriteHeader(204)
		return
	} else {
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(Domain.JsonError{401, 6})
		return
	}
}
