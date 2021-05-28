package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type Board struct {
	BoardId     string       `bson:"boardId" json:"boardId"`
	BoardName   string       `bson:"boardName" json:"boardName"`
	DomainName  string       `bson:"domainName" json:"domainName"`
	PostNum     int64        `bson:"postNum" json:"postNum"`
	Img         string       `bson:"img" json:"img"`
	ChildBoards []ChildBoard `bson:"childBoards" json:"childBoards"`
}

func (a *Board) TableName() string {
	return "Board"
}

func (a *Board) ToQueryBson() bson.M {
	if a.BoardId != "" {
		queryObject := bson.M{
			"boardId": a.BoardId,
		}
		return queryObject
	} else {
		queryObject := bson.M{
			"domainName": a.DomainName,
		}
		return queryObject
	}
}

type ChildBoard struct {
	BoardId        string `bson:"boardId" json:"boardId"`
	ChildBoardId   string `bson:"childBoardId" json:"childBoardId"`
	ChildBoardName string `bson:"childBoardName" json:"childBoardName"`
	PostNum        int64  `bson:"postNum" json:"postNum"`
}

func (a *ChildBoard) TableName() string {
	return "ChildBoard"
}

func (a *ChildBoard) ToQueryBson() bson.M {
	if a.ChildBoardId != "" {
		queryObject := bson.M{
			"childBoardId": a.ChildBoardId,
		}
		return queryObject
	} else {
		queryObject := bson.M{
			"boardId": a.BoardId,
		}
		return queryObject
	}

}

type Vote struct {
	VoteId         string    `bson:"voteId" json:"voteId"`
	Launcher       string    `bson:"launcher" json:"launcher"`
	BoardName      string    `bson:"boardName" json:"boardName"`
	DomainName     string    `bson:"domainName" json:"domainName"`
	Img            string    `bson:"img" json:"img"`
	Agree          int64     `bson:"agree" json:"agree"`
	Disagree       int64     `bson:"disagree" json:"disagree"`
	LauncherInfo   User      `bson:"launcherInfo" json:"launcherInfo"`
	AgreedUsers    []string  `bson:"agreedUsers" json:"agreedUsers"`
	DisagreedUsers []string  `bson:"disagreedUsers" json:"disagreedUsers"`
	Reason         string    `bson:"reason" json:"reason"`
	Deadline       time.Time `bson:"deadline" json:"deadline"`
	Status         string    `bson:"status" json:"status"`
}

func (a *Vote) TableName() string {
	return "Vote"
}

func (a *Vote) ToQueryBson() bson.M {
	if a.VoteId != "" {

		queryObject := bson.M{
			"voteId": a.VoteId,
		}
		return queryObject
	} else {
		queryObject := bson.M{
			"launcher": a.Launcher,
		}
		return queryObject
	}
}
