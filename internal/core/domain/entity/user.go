package entity

import (
	"time"
)

type (
	Role string
)

type User struct {
	UserID        string    `json:"user_id"`
	Name      string    `json:"first_name"`
	Surname   string    `json:"last_name"`
	Email     string    `json:"email"`
	Phone 	  string 	`json:"phone"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
