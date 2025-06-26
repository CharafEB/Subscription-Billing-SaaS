package middleware

import (
	pb "github.com/microservic/proto/castmerservice"
)

type Application struct {
	Address string
}

type Server struct {
	pb.UnimplementedClientHandlingServiceServer
}
