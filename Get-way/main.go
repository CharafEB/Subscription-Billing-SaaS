package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	castmer "github.com/microservic/getway/controller/Castmer"
	"github.com/microservic/getway/controller/subscription"
	"github.com/microservic/getway/middleware"
	"github.com/microservic/getway/router"
)

func main() {
	//Open the env file
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
	// get hte port of the server
	add := ":" + os.Getenv("PORT")
	app := middleware.Application{
		Address: add,
	}

	//pass the Application server
	objCastmerHandler := &castmer.Application{Application: app}
	objSubscriptionHandler := &subscription.Application{Application: app}

	appRouter := &router.Application{
		Application: app,
	}

	objrouter := &router.GetWayController{
		CastmerHandler:     objCastmerHandler,
		SubscritionHandler: objSubscriptionHandler,
	}

	mux := objrouter.Mux()
	if err := appRouter.Run(mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
