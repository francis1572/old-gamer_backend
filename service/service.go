package service

import (
	"context"
	"log"

	"final_backend/models"

	"go.mongodb.org/mongo-driver/mongo"
)

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
