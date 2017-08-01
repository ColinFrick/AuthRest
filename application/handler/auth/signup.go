package AuthHandler

import (
	"net/http"
	"encoding/json"
	"github.com/NiciiA/AuthRest/application/service"
	"github.com/NiciiA/AuthRest/application/domain"
	"github.com/NiciiA/AuthRest/application/dao"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"fmt"
	"github.com/NiciiA/AuthRest/config"
	"time"
	"strings"
)

type RegisterBody struct {
	PhoneNumber string `json:"phonenumber"`
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

func HandleSignup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	b := RegisterBody{}
	json.NewDecoder(r.Body).Decode(&b)

	b.FirstName = strings.TrimSpace(b.FirstName)
	b.LastName = strings.TrimSpace(b.LastName)

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

	u := Domain.User{}
	Dao.GetUsersCollection().Find(bson.M{"phonenumber": b.PhoneNumber}).One(&u)

	if u.ID.Hex() != "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(Domain.JsonError{400, 3})
		return
	}

	u.ID = bson.NewObjectId()
	u.PhoneNumber = b.PhoneNumber
	u.FirstName = b.FirstName
	u.LastName = b.LastName
	u.Role = "customer"
	u.Disabled = false
	u.CreatedDate = time.Now()
	u.UpdatedDate = time.Now()

	Dao.GetUsersCollection().Insert(&u)

	s := Domain.Settings{}
	s.User = u.ID
	s.Language = Config.DefaultLanguage

	Dao.GetSettingsCollection().Insert(&s)

	randomDigits := Service.Random(100000, 999999)
	er := Dao.GetRedisClient().Set(strconv.Itoa(randomDigits), u.ID.Hex(), 0).Err()
	if er != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(Domain.JsonError{500, 500})
		return
	}

	//TODO: Send SMS to c.PhoneNumber -> Content is randomDigits (sendSMS(c.PhoneNumber, randomDigits))
	fmt.Printf(strconv.Itoa(randomDigits))

	w.WriteHeader(201)
	return
}