package controller

import (
	"context"

	prt "github.com/microservic/proto/castmerservice"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) SignUp(ctx context.Context, Req *prt.SingUpData) (*prt.SingUpResponse, error) {
	if Req.Email == "" || Req.Password == "" {
		return &prt.SingUpResponse{Response: false}, status.Error(codes.InvalidArgument, "Email and password are required")
	}
	return &prt.SingUpResponse{Response: true}, nil
}
