package types

type SubscribtionData struct {
	UserName     string `json:"user_name"`
	UserLastname string `json:"user_lastName"`
	Email        string `json:"user_email"`
	UserPlan     string `json:"user_paln"`
	DayStart     string `json:"start_day"`
	DayEnd       string `json:"end_day"`
}
