package queue

import (
	"context"
	"fmt"
	"log"

	"github.com/microservic/subscription/middleware"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RApplication struct {
	Application *middleware.Application
	Chanl       *amqp.Channel
}

func (app *RApplication) StartUpdateListener(ctx context.Context) error {
	q, err := app.Chanl.QueueDeclare(
		"UpdateUser",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("queue declare error: %w", err)
	}

	msgs, err := app.Chanl.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("consume error: %w", err)
	}

	go func() {
		for {
			select {
			case d, ok := <-msgs:
				if !ok {
					log.Println("message channel closed")
					return
				}
				log.Printf("Received a message: %s", d.Body)
				app.Application.Store.Update.UpdateUser(d.Body)
			case <-ctx.Done():
				log.Println("listener stopped by context cancel")
				return
			}
		}
	}()
	//{"user_name":"salma","user_lastName":"benyamina","user_email":"fakefr123123@gmail.com"}
	return nil
}
