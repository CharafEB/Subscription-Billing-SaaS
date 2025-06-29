package model

import (
	"context"
	"fmt"

	"github.com/microservic/castmerservice/types"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
)

func (d *Data) Login(ctx context.Context, loginData types.LoginData) error {
	coll := d.DB.Database("Castmer").Collection("Singup")

	var res types.LoginData
	err := coll.FindOne(ctx, bson.M{"email": loginData.Email}).Decode(&res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("Email not found")
		}
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(loginData.Password))
	if err != nil {
		return fmt.Errorf("Wrong password")
	}

	res.Password = ""
	return nil
}
