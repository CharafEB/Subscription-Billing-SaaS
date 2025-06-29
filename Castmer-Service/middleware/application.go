package middleware

import (
	"github.com/microservic/castmerservice/model"
	pb "github.com/microservic/proto/castmerservice"
)

type Application struct {
	Address string
	Store model.Store
}

type Server struct {
	pb.UnimplementedClientHandlingServiceServer
}
