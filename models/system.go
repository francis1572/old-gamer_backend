package models

import "go.mongodb.org/mongo-driver/bson"

type System struct {
	TotalVotes    int   `bson:"totalVotes" json:"totalVotes"`
	VoteThreshold int64 `bson:"voteThreshold" json:"voteThreshold"`
	TotalBoards   int   `bson:"totalBoards" json:"totalBoards"`
}

type Notification struct {
	UserId     string `bson:"userId" json:"userId"`
	NotifyType string `bson:"notifyType" json:"notifyType"`
	Author     string `bson:"author" json:"author"`
	AuthorName string `bson:"authorName" json:"authorName"`
	PostId     string `bson:"postId" json:"postId"`
	Floor      int64  `bson:"floor" json:"floor"`
	VoteId     string `bson:"voteId" json:"voteId"`
}

func (a *Notification) TableName() string {
	return "Notification"
}

func (a *Notification) ToQueryBson() bson.M {
	queryObject := bson.M{
		"userId": a.UserId,
	}
	return queryObject
}
