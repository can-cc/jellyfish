package factory

import (
	"jellyfish/domain/user"
	"time"
)

func NewUser(id, username, hash, avatar string, createdAt, updatedAt time.Time) *user.AppUser {
	user := &user.AppUser{
		ID:        &id,
		Username:  &username,
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
	}
	user.SetAvatar(avatar)
	user.SetPasswordHash(hash)
	user.MarkSynced()
	return user
}
