package repository

import (
	"github.com/fwchen/jellyfish/domain/user"
)

type UserRepository interface {
	InsertUser(user *user.AppUser, hash string) error
	HasUser(username string) (bool, error)
	FindUser(username string) (*user.AppUser, error)
	FindUserHash(username string) (string, error)
}
