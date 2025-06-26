package router

import (
	castmer "github.com/microservic/getway/controller/Castmer"
	"github.com/microservic/getway/middleware"
)

type Application struct {
	middleware.Application
}

type CastmerHandler struct {
	CastmerHandler castmer.CastmerHandler
}

type GetWayController struct {
	CastmerHandler *castmer.Application
}
