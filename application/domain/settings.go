package Domain

import "gopkg.in/mgo.v2/bson"

type Settings struct {
	User bson.ObjectId `bson:"_id" json:"_id,omitempty"`
	Language string `bson:"language" json:"language,omitempty"`
}