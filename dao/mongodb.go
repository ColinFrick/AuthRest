package Dao

import (
	"gopkg.in/mgo.v2"
)

var session *mgo.Session

/**
	GetDatabase// projectaton
 */
func GetDatabase(s *mgo.Session) *mgo.Database {
	return s.DB("shelf_users")
}

/**
	Find collection in database
 */
func GetCollection(s *mgo.Session, collectionName string) *mgo.Collection {
	return GetDatabase(s).C(collectionName)
}

/**
	Find the cart collection
 */
func GetUsersCollection() *mgo.Collection {
	return GetCollection(session, "users")
}

/**
	init
 */
func init() {
	session, _ = mgo.Dial("mongodb://localhost")
	//GetItemCollection().Insert(Item{ID: bson.NewObjectId(), Barcode: "abc", Name: "Kekse", Image: "", Price: 2.50})
	//GetItemCollection().Insert(Item{ID: bson.NewObjectId(), Barcode: "abc2", Name: "Keksse", Image: "", Price: 2.50})
	//GetItemCollection().Insert(Item{ID: bson.NewObjectId(), Barcode: "abc3", Name: "Kekssse", Image: "", Price: 2.45})
}