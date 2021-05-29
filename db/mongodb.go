package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client

func DBConnection() {
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin:cr72tJVzEoSylRb7@sdmhw.ft9xv.mongodb.net/SDMHW?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
}

// getter for client
func GetDBCli() *mongo.Client {
	if client == nil {
		DBConnection()
	}
	// fmt.Println("client", client)
	return client
}
