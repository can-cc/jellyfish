package repository

import "jellyfish/domain/user"

type Repository interface {
	Save(user *user.AppUser) error
	Has(username string) (bool, error)
	FindByID(userID string) (*user.AppUser, error)
}
