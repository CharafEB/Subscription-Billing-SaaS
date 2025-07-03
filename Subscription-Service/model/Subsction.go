package model

import (
	"context"
	"github.com/microservic/subscription/types"
	"github.com/oklog/ulid/v2"
	"log"
)

func (d *Data) Subscribtion(ctx context.Context, signup types.SubscribtionData) error {
	subscriptionID := ulid.Make().String()

	qeury := `INSERT INTO  clients (sub_id,user_name , user_lastName ,user_paln ,start_day , end_day , user_email) VALUES ($1,$2,$3,$4,$5,$6,$7)`

	_, err := d.DB.ExecContext(ctx, qeury, &subscriptionID, &signup.UserName, &signup.UserLastname, &signup.UserPlan, &signup.DayStart, &signup.DayEnd, &signup.Email)
	if err != nil {
		log.Printf("Error adding event to display: %v", err)
		return err
	}

	return nil
}
