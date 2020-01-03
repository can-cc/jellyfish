package service

import "github.com/fwchen/jellyfish/domain/user"

type Guard interface {
	AuthenticateGuest(guest *user.Guest) error
	logGuestAccess(guest *user.Guest) error
	GenerateUserHash(guest *user.Guest) (string, error)
	compareGuestPassword(guest *user.Guest) error
}
