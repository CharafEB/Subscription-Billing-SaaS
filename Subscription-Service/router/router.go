package router

import (
	"log"
	"net"

	pb "github.com/microservic/proto/subscriptionservice"
	"github.com/microservic/subscription/controller"

	"google.golang.org/grpc"
)

func (app *Application) Run() error {

	log.Printf("Server is starting on Port %s", app.Address)

	lis, err := net.Listen("tcp", app.Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterClientHandlingServiceServer(s, &controller.Server{Application: &app.Application})
	log.Printf("server listening at %v", lis.Addr())
	s.Serve(lis)

	return nil
}
