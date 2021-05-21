package models

import "go.mongodb.org/mongo-driver/bson"

type Board struct {
	BoardId     string       `bson:"boardId" json:"boardId"`
	BoardName   string       `bson:"boardName" json:"boardName"`
	ChildBoards []ChildBoard `bson:"childBoards" json:"childBoards"`
}

func (a *Board) TableName() string {
	return "Board"
}

func (a *Board) ToQueryBson() bson.M {
	queryObject := bson.M{
		"boardId": a.BoardId,
	}
	return queryObject
}

type ChildBoard struct {
	ChildBoardId   string `bson:"childBoardId" json:"childBoardId"`
	ChildBoardName string `bson:"childBoardName" json:"childBoardName"`
	PostNum        int64  `bson:"postNum" json:"postNum"`
}

func (a *ChildBoard) TableName() string {
	return "ChildBoard"
}

func (a *ChildBoard) ToQueryBson() bson.M {
	queryObject := bson.M{
		"childBoardId": a.ChildBoardId,
	}
	return queryObject
}

type Vote struct {
	VoteId    string `bson:"voteId" json:"voteId"`
	BoardName string `bson:"boardName" json:"boardName"`
	Agree     int64  `bson:"agree" json:"agree"`
	Disagree  int64  `bson:"disagree" json:"disagree"`
}

func (a *Vote) TableName() string {
	return "Vote"
}

func (a *Vote) ToQueryBson() bson.M {
	queryObject := bson.M{
		"voteId": a.VoteId,
	}
	return queryObject
}
