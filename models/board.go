package models

type Board struct {
	BoardId     string       `bson:"boardId" json:"boardId"`
	BoardName   string       `bson:"boardName" json:"boardName"`
	ChildBoards []ChildBoard `bson:"childBoards" json:"childBoards"`
}

type ChildBoard struct {
	ChildBoardId   string `bson:"childBoardId" json:"childBoardId"`
	ChildBoardName string `bson:"childBoardName" json:"childBoardName"`
	PostNum        int64  `bson:"postNum" json:"postNum"`
}

type Vote struct {
	VoteId    string `bson:"voteId" json:"voteId"`
	BoardName string `bson:"boardName" json:"boardName"`
	Agree     int64  `bson:"agree" json:"agree"`
	Disagree  int64  `bson:"disagree" json:"disagree"`
}
