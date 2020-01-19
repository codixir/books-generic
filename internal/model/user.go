package model

import (
	"time"
)

type (
	//The Book Model
	User struct {
		ID        string    `json:"id"`
		Name      string    `json:"title"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
