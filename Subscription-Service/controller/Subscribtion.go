package controller

import (
	"context"

	prt "github.com/microservic/proto/subscriptionservice"
	"github.com/microservic/subscription/types"
)

func (s *Server) Subscribe(ctx context.Context, SubscribtionData *prt.SubscribtionData) (*prt.SubscribtionResponse, error) {
	Subscribtioninfo := types.SubscribtionData{
		UserName:     SubscribtionData.UserName,
		UserLastname: SubscribtionData.UserLastname,
		UserPlan:     SubscribtionData.UserPlan,
		DayStart:     SubscribtionData.DayStart,
		DayEnd:       SubscribtionData.DayEnd,
	}
	if err := s.Application.Store.Subscribtion.Subscribtion(ctx, Subscribtioninfo); err != nil {
		return &prt.SubscribtionResponse{Response: false	}, err
	}
	return &prt.SubscribtionResponse{Response: true}, nil
}
