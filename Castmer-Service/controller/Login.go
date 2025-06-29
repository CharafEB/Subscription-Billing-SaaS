package controller

import (
	"context"

	"github.com/microservic/castmerservice/types"
	prt "github.com/microservic/proto/castmerservice"
)

func (s *Server) Login(ctx context.Context, data *prt.LoginData) (*prt.LoginResponse, error) {
	LoginData := types.LoginData{
		Email:    data.Email,
		Password: data.Password,
	}
	if err := s.Application.Store.Login.Login(ctx, LoginData); err != nil {
		return &prt.LoginResponse{Response: false}, err
	}
	return &prt.LoginResponse{Response: true}, nil
}
