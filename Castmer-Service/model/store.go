package model

import (
	"context"

	"github.com/microservic/castmerservice/types"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Data struct {
	DB *mongo.Client
}

type Store struct {
	SingUp interface {
		SingUp(ctx context.Context, signup types.SingUpData) error
	}

	Login interface {
		Login(ctx context.Context, loginData types.LoginData) error
	}
}

func NewStore(db *mongo.Client) Store {
	if db == nil {
		panic("nil pointer passed to NewStore")
	}
	data := &Data{DB: db}
	return Store{
		SingUp: data,
		Login:  data,
	}
}
