package controller

import "fmt"

func (s *TrakerCron) Traker() error {
	s.Cron.AddFunc("@every 1s", func() { fmt.Print("Hello world") })
	return nil
}
