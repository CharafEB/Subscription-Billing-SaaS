package subscription

import (
	"github.com/microservic/getway/middleware"
	"net/http"
)

type Application struct {
	middleware.Application
}

type SubscritionHandler interface {
	Subscribtion(w http.ResponseWriter, r *http.Request)

}
