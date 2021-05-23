package service

import (
	"context"
	"log"
	"time"

	"final_backend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUser(db *mongo.Database, queryBson bson.M) (*models.User, error) {
	collection := db.Collection("GUser")
	var serviceResult = models.User{}
	cur := collection.FindOne(context.Background(), queryBson)
	// if no user then return nil
	if cur.Err() != nil {
		log.Println("Can't find user in DB")
		return nil, cur.Err()
	}
	// if has user then return
	err := cur.Decode(&serviceResult)
	if err != nil {
		log.Println("Decode user Error", err)
		return nil, err
	}
	log.Println("Get user:", serviceResult)
	return &serviceResult, nil
}

func GetUsers(db *mongo.Database) ([]models.User, error) {
	collection := db.Collection("GUser")
	var serviceResult = []models.User{}
	cur, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("Find user Error", err)
		return nil, err
	}
	for cur.Next(context.Background()) {
		result := models.User{}
		err := cur.Decode(&result)
		if err != nil {
			log.Println("Decode Article Error", err)
			return nil, err
		}
		serviceResult = append(serviceResult, result)
	}
	return serviceResult, nil
}

func SaveUser(db *mongo.Database, user models.User) (*mongo.InsertOneResult, error) {
	UserCollection := db.Collection("GUser")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := UserCollection.InsertOne(ctx, user)
	if err != nil {
		log.Println("Insert user Error", err)
		return nil, err
	}
	log.Println("Insert user Success", user)
	return res, nil
}

func GetBoardById(db *mongo.Database, query models.Board) (*models.Board, error) {
	BoardCollection := db.Collection("Board")
	var board models.Board
	result := BoardCollection.FindOne(context.Background(), query.ToQueryBson())
	err := result.Decode(&board)
	// log.Println("Success", board)
	if err != nil {
		log.Println("Decode task Error", err)
		return nil, err
	}
	return &board, nil
}

func GetBoardsByDomain(db *mongo.Database, task models.Board) ([]*models.Board, error) {
	BoardCollection := db.Collection("Board")
	var Boards []*models.Board

	cur, err := BoardCollection.Find(context.Background(), task.ToQueryBson())
	if err != nil {
		log.Println("Find answers Error", err)
		return nil, err
	}
	for cur.Next(context.Background()) {
		result := models.Board{}
		err := cur.Decode(&result)
		if err != nil {
			log.Println("Decode answer Error", err)
			return nil, err
		}
		Boards = append(Boards, &result)
	}
	return Boards, nil
}

func GetChildBoardByBoardId(db *mongo.Database, task models.ChildBoard) ([]*models.ChildBoard, error) {
	ChildBoardCollection := db.Collection("ChildBoard")
	var childBoards []*models.ChildBoard
	cur, err := ChildBoardCollection.Find(context.Background(), task.ToQueryBson())
	if err != nil {
		log.Println("Find answers Error", err)
		return nil, err
	}
	for cur.Next(context.Background()) {
		result := models.ChildBoard{}
		err := cur.Decode(&result)
		if err != nil {
			log.Println("Decode answer Error", err)
			return nil, err
		}
		childBoards = append(childBoards, &result)
	}
	return childBoards, nil
}

func GetPostsByBoardId(db *mongo.Database, task models.Board) ([]*models.Post, error) {
	PostCollection := db.Collection("Post")
	var posts []*models.Post
	cur, err := PostCollection.Find(context.Background(), task.ToQueryBson())
	if err != nil {
		log.Println("Find answers Error", err)
		return nil, err
	}
	for cur.Next(context.Background()) {
		result := models.Post{}
		err := cur.Decode(&result)
		if err != nil {
			log.Println("Decode answer Error", err)
			return nil, err
		}
		posts = append(posts, &result)
	}
	return posts, nil
}

func GetUserInfoById(db *mongo.Database, query models.User) (*models.User, error) {
	GUserCollection := db.Collection("GUser")
	var user models.User
	result := GUserCollection.FindOne(context.Background(), query.ToQueryBson())
	err := result.Decode(&user)
	// log.Println("Success", board)
	if err != nil {
		log.Println("Decode task Error", err)
		return nil, err
	}
	return &user, nil
}

func GetSpecialtyByUserId(db *mongo.Database, task models.Specialty) ([]*models.Specialty, error) {
	SpecialtyCollection := db.Collection("Specialty")
	var specialties []*models.Specialty
	cur, err := SpecialtyCollection.Find(context.Background(), task.ToQueryBson())
	if err != nil {
		log.Println("Find answers Error", err)
		return nil, err
	}
	for cur.Next(context.Background()) {
		result := models.Specialty{}
		err := cur.Decode(&result)
		if err != nil {
			log.Println("Decode answer Error", err)
			return nil, err
		}
		specialties = append(specialties, &result)
	}

	return specialties, nil
}

func GetPostsByUserId(db *mongo.Database, task models.Post) ([]*models.Post, error) {
	PostCollection := db.Collection("Post")
	var posts []*models.Post
	cur, err := PostCollection.Find(context.Background(), task.ToQueryBson())
	// log.Println("GetUserInfo posts:", task.ToQueryBson())
	if err != nil {
		log.Println("Find answers Error", err)
		return nil, err
	}
	for cur.Next(context.Background()) {
		result := models.Post{}
		err := cur.Decode(&result)
		if err != nil {
			log.Println("Decode answer Error", err)
			return nil, err
		}
		posts = append(posts, &result)
	}
	// log.Println("GetUserInfo posts:", posts)
	return posts, nil
}

func GetVotesByUserId(db *mongo.Database, task models.Vote) ([]*models.Vote, error) {
	VoteCollection := db.Collection("Vote")
	var votes []*models.Vote
	cur, err := VoteCollection.Find(context.Background(), task.ToQueryBson())
	// log.Println("GetUserInfo Votes:", task.ToQueryBson())
	if err != nil {
		log.Println("Find answers Error", err)
		return nil, err
	}
	for cur.Next(context.Background()) {
		result := models.Vote{}
		err := cur.Decode(&result)
		if err != nil {
			log.Println("Decode answer Error", err)
			return nil, err
		}
		votes = append(votes, &result)
	}
	// log.Println("GetUserInfo votes:", votes)
	return votes, nil
}

func GetVote(db *mongo.Database) ([]*models.Vote, error) {
	VoteCollection := db.Collection("Vote")
	var votes []*models.Vote
	cur, err := VoteCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("GetVote Error", err)
		return nil, err
	}

	for cur.Next(context.Background()) {
		result := models.Vote{}
		err := cur.Decode(&result)
		if err != nil {
			log.Println("Decode Vote Error", err)
			return nil, err
		}
		votes = append(votes, &result)
	}
	return votes, nil
}
