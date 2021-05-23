package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type Post struct {
	PostId       string     `bson:"postId" json:"postId"`
	BoardId      string     `bson:"boardId" json:"boardId"`
	ChildBoardId string     `bson:"childBoardId" json:"childBoardId"`
	PostTag      string     `bson:"postTag" json:"postTag"`
	PostTitle    string     `bson:"postTitle" json:"postTitle"`
	Author       string     `bson:"author" json:"author"`
	Content      []Block    `bson:"content" json:"content"`
	Floor        int64      `bson:"floor" json:"floor"`
	CommentNum   int64      `bson:"commentNum" json:"commentNum"`
	LikeNum      int64      `bson:"likeNum" json:"likeNum"`
	Time         time.Time  `bson:"time" json:"time"`
	Citations    []Citation `bson:"citation" json:"citation"`
	LikedUsers   []string   `bson:"likedUsers" json:"likedUsers"`
	AuthorInfo   User       `bson:"authorInfo" json:"authorInfo"`
}

func (a *Post) TableName() string {
	return "Post"
}

func (a *Post) ToQueryBson() bson.M {
	if a.PostId != "" {
		queryObject := bson.M{
			"postId": a.PostId,
		}
		return queryObject
	} else {
		queryObject := bson.M{
			"author": a.Author,
		}
		return queryObject
	}

}

type Block struct {
	PostId   string `bson:"postId" json:"postId"`
	Floor    int64  `bson:"floor" json:"floor"`
	BlockId  int64  `bson:"blockId" json:"blockId"`
	Subtitle string `bson:"subtitle" json:"subtitle"`
	Content  string `bson:"content" json:"content"`
}

func (a *Block) TableName() string {
	return "Block"
}

func (a *Block) ToQueryBson() bson.M {

	if a.PostId != "" {
		queryObject := bson.M{
			"postId": a.PostId,
			"floor":  a.Floor,
		}
		return queryObject
	} else {
		queryObject := bson.M{
			"blockId": a.BlockId,
		}
		return queryObject
	}

}

type Citation struct {
	CitationId string `bson:"citationId" json:"citationId"`
	PostId     string `bson:"postId" json:"postId"`
	Floor      int64  `bson:"floor" json:"floor"`
	CitedFloor int64  `bson:"citedFloor" json:"citedFloor"`
	BlockId    int64  `bson:"blockId" json:"blockId"`
}

func (c *Citation) TableName() string {
	return "Citation"
}

func (c *Citation) ToQueryBson() bson.M {
	if c.PostId != "" {
		queryObject := bson.M{
			"postId": c.PostId,
			"floor":  c.Floor,
		}
		return queryObject
	} else {
		queryObject := bson.M{
			"citationId": c.CitationId,
		}
		return queryObject
	}

}

type Comment struct {
	CommentId  string    `bson:"commentId" json:"commentId"`
	PostId     string    `bson:"postId" json:"postId"`
	Tag        string    `bson:"tag" json:"tag"`
	Floor      int64     `bson:"floor" json:"floor"`
	Content    string    `bson:"content" json:"content"`
	Author     User      `bson:"author" json:"author"`
	LikeNum    int64     `bson:"likeNum" json:"likeNum"`
	LikedUsers []User    `bson:"likedUsers" json:"likedUsers"`
	Time       time.Time `bson:"time" json:"time"`
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
