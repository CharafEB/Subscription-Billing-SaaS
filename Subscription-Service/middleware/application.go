package middleware

import (
	"github.com/microservic/subscription/model"
)

type Application struct {
	Address string
	Store   model.Store
}
