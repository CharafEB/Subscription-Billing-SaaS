package castmer

import (
	"github.com/microservic/getway/middleware"
	"net/http"
)

type Application struct {
	middleware.Application
}

type CastmerHandler interface {
	SingIn(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}
