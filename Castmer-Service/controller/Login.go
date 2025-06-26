package controller
import (
	"context"
	"fmt"

	prt "github.com/microservic/proto/castmerservice"
)

func (s *Server) Login(ctx context.Context, req *prt.LoginData) (*prt.LoginResponse, error) {
	fmt.Print("Hellow")
	out := &prt.LoginResponse{
		Response: true}
	return out, nil
}
