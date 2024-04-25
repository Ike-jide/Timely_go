package users

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Users model
type Users struct {
	ID       bson.ObjectId `json:"_id" bson:"_id"`
	UserName string        `json:"user_name" bson:"user_name"`
	Email    string        `json:"email" bson:"email"`
	Date     time.Time     `json:"date" bson:"date"`
}
