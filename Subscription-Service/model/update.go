package model

import (
	"encoding/json"
	"fmt"

	"github.com/microservic/subscription/types"
)

func (d *Data) UpdateUser(msg []byte) error {
	var input types.TrakerResponse

	err := json.Unmarshal(msg, &input)
	if err != nil {
		return fmt.Errorf("failed to parse user data: %v", err)
	}
	query := `UPDATE clients SET is_notify = true WHERE user_email = $1`

	res, err := d.DB.Exec(query, input.Email)
	if err != nil {
		return fmt.Errorf("failed to execute update: %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to retrieve affected rows: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no user found with email: %s", input.Email)
	}

	return nil
}
