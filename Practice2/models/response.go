package models

type UserResponse struct {
	UserID int `json:"user_id"`
}

type CreateUserResponse struct {
	Created string `json:"created"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
