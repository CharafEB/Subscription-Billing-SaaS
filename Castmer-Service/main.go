package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/microservic/castmerservice/middleware"
	"github.com/microservic/castmerservice/model"
	"github.com/microservic/castmerservice/router"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
	add := ":" + os.Getenv("PORT")
	log.Println("Store has been opened")

	uri := os.Getenv("MONGODB_URI")
	client, err := mongo.Connect(options.Client().
		ApplyURI(uri))

	if err != nil {
		log.Panic(err)
	}
	log.Println("Database connection successful")

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	store := model.NewStore(client)
	app := middleware.Application{
		Address: add,
		Store:   store,
	}

	objRouter := &router.Application{
		Application: app,
	}

	if err := objRouter.Run(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
