package models

import (
	"go.mongodb.org/mongo-driver/bson"
)

//User structure
type User struct {
	Name                  string      `bson:"name" json:"name"`
	AccessToken           string      `bson:"accessToken" json:"accessToken"`
	ImageUrl              string      `bson:"imageUrl" json:"imageUrl"`
	Email                 string      `bson:"email" json:"email"`
	FamilyName            string      `bson:"familyName" json:"familyName"`
	GivenName             string      `bson:"givenName" json:"givenName"`
	UserId                string      `bson:"userId" json:"userId"`
	SelfIntroduction      string      `bson:"selfIntroduction" json:"selfIntroduction"`
	InterestGames         string      `bson:"interestGames" json:"interestGames"`
	CumulateGameSpecialty []Specialty `bson:"cumulateGameSpecialty" json:"cumulateGameSpecialty"`
	PublishPost           []Post      `bson:"publishPost" json:"publishPost"`
	LaunchNewBoard        []Vote      `bson:"launchNewBoard" json:"launchNewBoard"`
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

type Specialty struct {
	BoardId   string `bson:"boardId" json:"boardId"`
	BoardName string `bson:"boardName" json:"boardName"`
	UserId    string `bson:"userId" json:"userId"`
	Score     int64  `bson:"score" json:"score"`
}

func (s *Specialty) ToQueryBson() bson.M {
	var queryObject bson.M
	if s.UserId != "" {
		queryObject = bson.M{
			"userId": s.UserId,
		}
	} else {
		queryObject = bson.M{
			"userId": s.UserId,
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
