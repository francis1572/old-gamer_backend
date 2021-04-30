package models

import (
	"go.mongodb.org/mongo-driver/bson"
)

//User structure
type Test struct {
	UserId string `bson:"userId" json:"userId"`
	Title  string `bson:"title" json:"title"`
	Text   string `bson:"text" json:"text"`
}

//TableName return name of database table
func (a *Test) TableName() string {
	return "Test"
}

func (a *Test) ToQueryBson() bson.M {
	queryObject := bson.M{
		"userId": a.UserId,
	}
	return queryObject
}
