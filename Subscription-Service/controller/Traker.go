package controller

import (
	"context"
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

func (s *TrakerCron) Traker(ctx context.Context) error {
	s.Cron.AddFunc("@every 24h", func() {
		Users, err := s.Application.Store.Traker.Traker(ctx)
		if err != nil || len(Users) == 0 {
			fmt.Errorf("Failed to get users: %v", err)
			return
		}

		q, err := s.Application.Chanl.QueueDeclare(
			"NotifyUser",
			false,
			false,
			false,
			false,
			nil,
		)

		if err != nil {
			fmt.Print("Failed to declare a queue")
			fmt.Errorf(err.Error())
		}
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		msg, _ := json.Marshal(Users)
		err = s.Application.Chanl.PublishWithContext(ctx,
			"",
			q.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        msg,
			})

		if err != nil {
			fmt.Print("Failed to publish a message")
			fmt.Errorf(err.Error())
		}

	})

	return nil
}
