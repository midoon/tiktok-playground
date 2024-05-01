package dto

type UserData struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type LoginResponse struct {
	Status  bool     `json:"status"`
	Message string   `json:"message"`
	Data    UserData `json:"data"`
}
