package models

import (
	"fmt"
	"github.com/eugenebogun/music-storage/components"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// User model.
type User struct {
	Id       bson.ObjectId `json:"id",bson:"_id,omitempty"`
	Email    string        `json:"email",bson:"email"`
	Password string        `json:"password",bson:"password"`
}

var db *mgo.Session
var table = "users"

// Insert user data into user table.
func (user *User) Insert() error {
	var dbErr error
	db, dbErr = mgo.Dial("localhost")
	defer db.Close()

	if dbErr != nil {
		return dbErr
	}

	c := db.DB(components.DbName).C(table)
	count, insertError := c.Find(bson.M{"Email": user.Email}).Limit(1).Count()
	if insertError != nil {
		return insertError
	}
	if count > 0 {
		return fmt.Errorf("User %s already exists", user.Email)
	}
	return c.Insert(user)
}
