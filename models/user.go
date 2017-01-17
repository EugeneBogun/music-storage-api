package models

import (
    "gopkg.in/mgo.v2/bson"
    "gopkg.in/mgo.v2"
    "fmt"
    "github.com/eugenebogun/music-storage/components"
)

type User struct {
    Id       bson.ObjectId `json:"id",bson:"_id,omitempty"`
    Email    string `json:"email",bson:"email"`
    Password string `json:"password",bson:"password"`
}

var db *mgo.Session
var table = "users"

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