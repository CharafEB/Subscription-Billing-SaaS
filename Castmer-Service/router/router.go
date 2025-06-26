package router

import (
	"log"
	"net"

	"github.com/microservic/castmerservice/controller"
	pb "github.com/microservic/proto/castmerservice"

	"google.golang.org/grpc"
)

func (app *Application) Run() error {

	log.Printf("Server is starting on Port %s", app.Address)

	lis, err := net.Listen("tcp", app.Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterClientHandlingServiceServer(s, &controller.Server{})
	log.Printf("server listening at %v", lis.Addr())
	s.Serve(lis)

	return nil
}
