package service

import (
	"github.com/fwchen/jellyfish/domain/visitor"
)

type Guard interface {
	AuthenticateGuest(guest *visitor.Visitor) error
	logGuestAccess(guest *visitor.Visitor) error
	GenerateUserHash(guest *visitor.Visitor) (string, error)
	compareGuestPassword(guest *visitor.Visitor) error
}
