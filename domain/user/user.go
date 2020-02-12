package user

import (
	"time"
)

type AppUser struct {
	ID           *string
	Username     *string
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
	passwordHash *string
	avatar       *Avatar
	synced       bool
}

func (a *AppUser) IsSynced() bool {
	return a.synced
}

func (a *AppUser) MarkSynced() {
	a.synced = true
}

func (a *AppUser) GetAvatar() *Avatar {
	return a.avatar
}

func (a *AppUser) SetAvatar(fileName string) {
	if fileName == "" {
		a.avatar = nil
	} else {
		a.avatar = &Avatar{FileName: fileName}
	}

}

func (a *AppUser) GetPasswordHash() string {
	return *a.passwordHash
}

func (a *AppUser) SetPasswordHash(hash string) {
	a.passwordHash = &hash
}
