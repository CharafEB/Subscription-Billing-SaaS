package middleware

import (
	"github.com/microservic/subscription/model"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Application struct {
	Address string
	Store   model.Store
	Chanl   *amqp.Channel
}
