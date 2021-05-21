package models

import "time"

type Post struct {
	PostId       string    `bson:"postId" json:"postId"`
	BoardId      string    `bson:"boardId" json:"boardId"`
	ChildBoardId string    `bson:"childBoardId" json:"childBoardId"`
	Content      string    `bson:"content" json:"content"`
	CommentNum   int64     `bson:"commentNum" json:"commentNum"`
	LikeNum      int64     `bson:"likeNum" json:"likeNum"`
	Time         time.Time `bson:"time" json:"time"`
	ReplyPosts   []Post    `bson:"replyPosts" json:"replyPosts"`
}

type Comment struct {
	CommentId string    `bson:"commentId" json:"commentId"`
	PostId    string    `bson:"postId" json:"postId"`
	Content   string    `bson:"content" json:"content"`
	LikeNum   int64     `bson:"likeNum" json:"likeNum"`
	Time      time.Time `bson:"time" json:"time"`
}
