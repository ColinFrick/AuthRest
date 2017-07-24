package AuthHandler

import (
	"net/http"
	"github.com/NiciiA/AuthRest/application/service"
	"encoding/json"
	"github.com/NiciiA/AuthRest/application/domain"
	"github.com/NiciiA/AuthRest/application/dao"
	"gopkg.in/mgo.v2/bson"
	"github.com/NiciiA/AuthRest/application/jwt"
)

type VerifyResponse struct {
	User Domain.User `json:"user"`
	Token string `json:"token"`
}

func HandleVerify(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	c := r.URL.Query().Get("c")

	reg, err := Service.RanValidator(c)
	if !reg || err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(Domain.JsonError{400, 5})
		return
	}

	val, err := Dao.GetRedisClient().Get(c).Result()
	if val == "" || err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(Domain.JsonError{400, 5})
		return
	}
	Dao.GetRedisClient().Del(c)

	u := Domain.User{}
	Dao.GetUsersCollection().Find(bson.M{"_id": bson.ObjectIdHex(val)}).One(&u)

	if u.ID.Hex() == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(Domain.JsonError{400, 2})
		return
	}

	token := JWT.CreateToken(u)
	json.NewEncoder(w).Encode(VerifyResponse{u, token})

	w.WriteHeader(200)
	return
}