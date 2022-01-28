package service

import (
	"jellyfish/domain/visitor"
)

type Guard interface {
	Authenticate(guest *visitor.Visitor, hash string)
	//logGuestAccess(guest *visitor.Visitor) error
	GeneratePasswordHash(guest *visitor.Visitor) (string, error)
	//compareGuestPassword(guest *visitor.Visitor, hash string) error
}
