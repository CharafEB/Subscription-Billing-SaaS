package controller

import (
	pb "github.com/microservic/proto/subscriptionservice"
	"github.com/microservic/subscription/middleware"
	"github.com/robfig/cron/v3"
)

type Application struct {
	*middleware.Application
}
type Server struct {
	pb.UnimplementedClientHandlingServiceServer
	Application *middleware.Application
}

type TrakerCron struct {
	Cron        *cron.Cron
	Application *middleware.Application
}
