package model

import (
	"context"
	"database/sql"

	"github.com/microservic/subscription/types"
)

type Data struct {
	DB *sql.DB
}

type Store struct {
	Subscribtion interface {
		Subscribtion(ctx context.Context, signup types.SubscribtionData) error
	}

	Traker interface {
		Traker(ctx context.Context) ([]types.TrakerResponse, error)
	}
	Update interface {
		UpdateUser(msg []byte) error
	}
}

func NewStore(db *sql.DB) Store {
	if db == nil {
		panic("nil pointer passed to NewStore")
	}
	data := &Data{DB: db}
	return Store{
		Subscribtion: data,
		Traker:       data,
		Update:       data,
	}
}
