package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/CesarDelgadoM/tutorials-api/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongodb *mongo.Database

func ConnectMongoDB() *mongo.Database {

	if mongodb == nil {
		config.InitConfig()

		name := os.Getenv("NAME_DB")
		user := os.Getenv("USER_DB")
		password := os.Getenv("PASSWORD_DB")

		stringConn := fmt.Sprintf("mongodb+srv://%s:%s@dev-go.jv7pq.mongodb.net/?retryWrites=true&w=majority", user, password)

		clientOptions := options.Client().ApplyURI(stringConn)

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatal(err)
		}

		err = client.Ping(ctx, nil)
		if err != nil {
			log.Fatal(err)
		}

		mongodb = client.Database(name)
	}

	fmt.Println("Connected to MongoDB!")
	return mongodb
}
