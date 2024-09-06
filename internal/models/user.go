package models

import "time"

type User struct {
	ID          int64     `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Password    string    `json:"-"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	DateOfBirth time.Time `json:"date_of_birth"`
	CreatedAt    time.Time `json:"created_at"`
	LastLogin   time.Time `json:"last_login"`
}
