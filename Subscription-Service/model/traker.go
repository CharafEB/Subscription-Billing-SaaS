package model

import (
	"context"

	"time"

	"github.com/microservic/subscription/types"
)

func (d *Data) Traker(ctx context.Context) ([]types.TrakerResponse, error) {
	targetDate := time.Now().AddDate(0, 0, 3).Format("2006-01-02")
	query := `
		SELECT user_name, user_lastName, user_email
		FROM clients 
		WHERE end_day::date = $1 AND is_notify = false
	`

	rows, err := d.DB.QueryContext(ctx, query, targetDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var trackerResponses []types.TrakerResponse

	for rows.Next() {
		var tr types.TrakerResponse
		err := rows.Scan(&tr.UserName, &tr.UserLastname, &tr.Email)
		if err != nil {
			continue
		}

		trackerResponses = append(trackerResponses, tr)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return trackerResponses, nil
}
