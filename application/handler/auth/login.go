package AuthHandler

import (
	"net/http"
	"encoding/json"
	"github.com/NiciiA/AuthRest/application/domain"
	"github.com/NiciiA/AuthRest/application/service"
	"gopkg.in/mgo.v2/bson"
	"github.com/NiciiA/AuthRest/application/dao"
	"strconv"
	"fmt"
)

type LoginBody struct {
	PhoneNumber string `json:"phonenumber"`
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	b := LoginBody{}
	json.NewDecoder(r.Body).Decode(&b)

	reg, err := Service.PhoneNumberValidator(b.PhoneNumber)
	if !reg || err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(Domain.JsonError{400, 1})
		return
	}

	u := Domain.User{}
	Dao.GetUsersCollection().Find(bson.M{"phonenumber": b.PhoneNumber, "disabled": false}).One(&u)

	if u.ID.Hex() == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(Domain.JsonError{400, 2})
		return
	}

	randomDigits := Service.Random(100000, 999999)
	er := Dao.GetRedisClient().Set(strconv.Itoa(randomDigits), u.ID.Hex(), 0).Err()
	if er != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(Domain.JsonError{500, 500})
		return
	}

	//TODO: Send SMS to c.PhoneNumber -> Content is randomDigits (sendSMS(c.PhoneNumber, randomDigits))
	fmt.Printf(strconv.Itoa(randomDigits))

	w.WriteHeader(200)
	return
}

