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

func GetTest(db *mongo.Database, query models.Test) ([]models.Test, error) {
	collection := db.Collection("test")

	var tests []models.Test

	cur, err := collection.Find(context.Background(), query.ToQueryBson())
	if err != nil {
		log.Println("Find Articles Error", err)
		return nil, err
	}
	for cur.Next(context.Background()) {
		result := models.Test{}
		err := cur.Decode(&result)
		if err != nil {
			log.Println("Decode Article Error", err)
			return nil, err
		}

		tests = append(tests, result)
	}
	return tests, nil
}
