package model

import "time"

type Ad struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title" validate:"required,min=5,max=100"`
	Description string    `json:"description" validate:"required,min=10,max=1000"`
	ImageURL    string    `json:"image_url" validate:"required,url"`
	Price       uint      `json:"price" validate:"required,gt=0,lt=2147483647"`
	CreatedAt   time.Time `json:"created_at"`
	UserID      uint      `json:"-"`
	User        User      `json:"-"`
}

type AdNewRequest struct {
	Title       string `json:"title" validate:"required,min=5,max=100"`
	Description string `json:"description" validate:"required,min=10,max=1000"`
	ImageURL    string `json:"image_url" validate:"required,url"`
	Price       uint   `json:"price" validate:"required,gt=0"`
}

type AdResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url"`
	Price       uint      `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	Author      string    `json:"author"`
	IsMine      bool      `json:"is_mine,omitempty"`
}
