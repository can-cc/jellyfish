package user

import (
	"time"
)

type AppUser struct {
	ID        string     `json:"id"`
	Username  string     `json:"username"`
	AvatarID  *string    `json:"avatarID"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}
