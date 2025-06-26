package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/microservic/castmerservice/middleware"
	"github.com/microservic/castmerservice/router"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
	add := ":" + os.Getenv("PORT")
	log.Println("Store has been opened")

	app := middleware.Application{
		Address: add,
	}

	objRouter := &router.Application{
		Application: app,
	}

	if err := objRouter.Run(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
