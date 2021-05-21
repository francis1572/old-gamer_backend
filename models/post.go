package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type Post struct {
	PostId       string    `bson:"postId" json:"postId"`
	Floor        int64     `bson:"floor" json:"floor"`
	BoardId      string    `bson:"boardId" json:"boardId"`
	ChildBoardId string    `bson:"childBoardId" json:"childBoardId"`
	Content      string    `bson:"content" json:"content"`
	CommentNum   int64     `bson:"commentNum" json:"commentNum"`
	LikeNum      int64     `bson:"likeNum" json:"likeNum"`
	Time         time.Time `bson:"time" json:"time"`
	ReplyPosts   []Post    `bson:"replyPosts" json:"replyPosts"`
}

func (a *Post) TableName() string {
	return "Post"
}

func (a *Post) ToQueryBson() bson.M {
	queryObject := bson.M{
		"postId": a.PostId,
	}
	return queryObject
}

type Comment struct {
	CommentId string    `bson:"commentId" json:"commentId"`
	PostId    string    `bson:"postId" json:"postId"`
	Floor     int64     `bson:"floor" json:"floor"`
	Content   string    `bson:"content" json:"content"`
	LikeNum   int64     `bson:"likeNum" json:"likeNum"`
	Time      time.Time `bson:"time" json:"time"`
}

func (a *Comment) TableName() string {
	return "Comment"
}

func (a *Comment) ToQueryBson() bson.M {
	queryObject := bson.M{
		"commentId": a.CommentId,
	}
	return queryObject
}
