package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	castmer "github.com/microservic/getway/controller/Castmer"
	"github.com/microservic/getway/middleware"
	"github.com/microservic/getway/router"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	add := ":" + os.Getenv("PORT")
	app := middleware.Application{
		Address: add,
	}

	objCastmerHandler := &castmer.Application{Application: app}
	appRouter := &router.Application{Application: app}
	objrouter := &router.GetWayController{CastmerHandler: objCastmerHandler}
	mux := objrouter.Mux()
	if err := appRouter.Run(mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
