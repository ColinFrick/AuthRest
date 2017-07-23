package Domain

import "gopkg.in/mgo.v2/bson"

type Customer struct {
	ID bson.ObjectId `bson:"_id" json:"id,omitempty"`
	CardNumber string `bson:"cardnumber" json:"cardnumber,omitempty"`
	PhoneNumber string `bson:"phonenumber" json:"phonenumber,omitempty"`
	FirstName string `bson:"firstname" json:"firstname,omitempty"`
	LastName string `bson:"lastname" json:"lastname,omitempty"`
	Disabled bool `bson:"disabled" json:"disabled,omitempty"`
}

type Customers []Customer
