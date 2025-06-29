package controller

import (
	"context"
	"fmt"

	"github.com/microservic/castmerservice/types"
	prt "github.com/microservic/proto/castmerservice"
)

func (s *Server) SignUp(ctx context.Context, data *prt.SingUpData) (*prt.SingUpResponse, error) {
	pass, _ := hashPassword(data.Password)
	user := types.SingUpData{
		ID:        data.Id,
		FirstName: data.Name,
		LastName:  data.Lasname,
		Email:     data.Email,
		Password:  pass,
	}
	fmt.Print(user)
	err := s.Application.Store.SingUp.SingUp(ctx, user)
	if err != nil {
		return &prt.SingUpResponse{Response: false}, err
	}
	return &prt.SingUpResponse{Response: true}, nil
}
