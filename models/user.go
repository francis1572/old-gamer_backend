package models

type User struct {
	UserId      string `bson:"userId" json:"userId"`
	Name        string `bson:"name" json:"name"`
	AccessToken string `bson:"accessToken" json:"accessToken"`
	ImageUrl    string `bson:"imageUrl" json:"imageUrl"`
	Email       string `bson:"email" json:"email"`
	FamilyName  string `bson:"familyName" json:"familyName"`
	GivenName   string `bson:"givenName" json:"givenName"`
}
