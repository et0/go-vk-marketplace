package model

import "time"

type User struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username" validate:"required,min=3,max=32"`
	Password  string    `json:"password" validate:"required,min=8,max=32"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required,min=3,max=32"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type SignupRequest struct {
	Username string `json:"username" validate:"required,min=3,max=32"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}
