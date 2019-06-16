package models

import (
	"time"
)

// User :
type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"createdAt"`
}
