package controller

import (
	"github.com/microservic/castmerservice/middleware"
	pb "github.com/microservic/proto/castmerservice"
)

type Application struct {
	*middleware.Application
}
type Server struct {
	pb.UnimplementedClientHandlingServiceServer
	Application *middleware.Application
}
