package types

type SingUpData struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// example: {"user_id":"12345", "user_name":"john_doe", "last_name":"Doe", "email":"
