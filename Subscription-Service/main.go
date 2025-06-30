package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/microservic/subscription/controller"
	"github.com/microservic/subscription/middleware"
	"github.com/microservic/subscription/model"
	"github.com/microservic/subscription/router"
	"github.com/robfig/cron/v3"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
	add := ":" + os.Getenv("PORT")
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
	// start connacting the data base of the subscription service
	store := model.NewStore(db)
	app := middleware.Application{
		Address: add,
		Store:   store,
	}
	// runing the traker by calling the cron
	Traker := controller.TrakerCron{
		Cron: cron.New(),
	}
	Traker.Traker()
	//Traker.Cron.Start()

	//passing the data for the router so that we can run it
	objRouter := &router.Application{
		Application: app,
	}

	if err := objRouter.Run(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
