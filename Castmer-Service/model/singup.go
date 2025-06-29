package model

import (
	"context"
	"fmt"

	"github.com/microservic/castmerservice/types"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (d *Data) SingUp(ctx context.Context, signup types.SingUpData) error {
	coll := d.DB.Database("Castmer").Collection("Singup")
	_, err := coll.InsertOne(context.TODO(), signup)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return fmt.Errorf("this email is linked pless try new one")
		}
		return err
	}

	return nil
}
