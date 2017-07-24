package Domain

import "gopkg.in/mgo.v2/bson"

type User struct {
	ID bson.ObjectId `bson:"_id" json:"_id,omitempty"`
	//CardNumber string `bson:"cardnumber" json:"cardnumber,omitempty"`
	PhoneNumber string `bson:"phonenumber" json:"phonenumber,omitempty"`
	FirstName string `bson:"firstname" json:"firstname,omitempty"`
	LastName string `bson:"lastname" json:"lastname,omitempty"`
	Role string `bson:"role" json:"role,omitempty"`
	Disabled bool `bson:"disabled" json:"disabled"`
}

type Users []User
