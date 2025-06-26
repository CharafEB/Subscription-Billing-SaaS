package types

type SingIn struct {
	UserId   string `json:"user_id"`
	UserName string `json:"user_name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// example: {"user_id":"12345", "user_name":"john_doe", "last_name":"Doe", "email":"