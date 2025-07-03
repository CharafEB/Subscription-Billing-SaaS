package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/microservic/subscription/controller"
	"github.com/microservic/subscription/middleware"
	"github.com/microservic/subscription/model"
	"github.com/microservic/subscription/queue"
	"github.com/microservic/subscription/router"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/robfig/cron/v3"
)

func main() {
	//open the env file
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
	//get the server port
	add := ":" + os.Getenv("PORT")

	//get the database link
	link := os.Getenv("Database_Link")
	db, err := sql.Open("postgres", link)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("Database connection successful")

	conn, err := amqp.Dial(os.Getenv("RabitMQ_Link"))
	if err != nil {
		fmt.Print("Failed to connect to RabbitMQ")
		log.Fatal(err)
	}
	defer conn.Close()
	log.Println("connect to RabbitMQ successful ")
	ch, err := conn.Channel()
	if err != nil {
		fmt.Print("Failed to open a channel")
		log.Fatal(err)
	}
	defer ch.Close()
	log.Println("open a channel successful ")

	// start connacting the data base of the subscription service
	store := model.NewStore(db)
	app := middleware.Application{
		Address: add,
		Store:   store,
		Chanl:   ch,
	}
	ctx := context.Background()

	rapp := &queue.RApplication{
		Application: &app,
		Chanl:       ch,
	}
	if err := rapp.StartUpdateListener(ctx); err != nil {
		log.Fatalf("failed to start listener: %v", err)
	}
	// runing the traker by calling the cron
	Traker := controller.TrakerCron{
		Cron:        cron.New(),
		Application: &app,
	}

	Traker.Traker(ctx)
	Traker.Cron.Start()

	//passing the data for the router so that we can run it
	objRouter := &router.Application{
		Application: app,
	}

	if err := objRouter.Run(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
