package repository

import "github.com/fwchen/jellyfish/domain/user"

type Repository interface {
	//InsertUser(user *AppUser, hash string) error
	Save(user *user.AppUser) error
	Has(username string) (bool, error)
	FindByID(userID string) (*user.AppUser, error)
	//FindUserHashByName(username string) (string, error)
	//UpdateUserAvatar(user *AppUser, avatar string) error
	//FindUerAvatar(userID string) (*Avatar, error)
}
