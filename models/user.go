package models

import (
	"go.mongodb.org/mongo-driver/bson"
)

//User structure
type User struct {
	Name        string `bson:"name" json:"name"`
	AccessToken string `bson:"accessToken" json:"accessToken"`
	ImageUrl    string `bson:"imageUrl" json:"imageUrl"`
	Email       string `bson:"email" json:"email"`
	FamilyName  string `bson:"familyName" json:"familyName"`
	GivenName   string `bson:"givenName" json:"givenName"`
	UserId      string `bson:"userId" json:"userId"`
}

func (u *User) ToQueryBson() bson.M {
	var queryObject bson.M
	if u.UserId != "" {
		queryObject = bson.M{
			"userId": u.UserId,
		}
	} else {
		queryObject = bson.M{
			"email": u.Email,
		}
	}
	return queryObject
}

type Auth struct {
	ProjectId  int    `bson:"projectId" json:"projectId"`
	UserId     string `bson:"userId" json:"userId"`
	CodeType   string `bson:"codeType" json:"codeType"`
	StatusCode string `bson:"statusCode" json:"statusCode"`
}

type Auths []Auth

//TableName return name of database table
func (u *Auth) TableName() string {
	return "Authentication"
}

func (u *Auth) ToQueryBson() bson.M {
	queryObject := bson.M{
		"userId":     u.UserId,
		"statusCode": u.StatusCode,
	}
	return queryObject
}

func (a Auths) SelectProjectIdList() []int {
	var list []int
	for _, user := range a {
		list = append(list, user.ProjectId)
	}
	return list
}

type Success struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

var (
	InsertSuccess = Success{Success: true, Message: "insert success"}
)
