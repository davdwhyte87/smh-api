package models

import (
	"gopkg.in/mgo.v2/bson"
)

// Trip ... this is a representation of a trip on the server
type Trip struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string      `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`
	Quantity int `bson:"amount" json:"amount"`
	Price int `bson:"price" json:"price"`
}
