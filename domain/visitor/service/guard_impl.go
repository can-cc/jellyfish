package service

import (
	"github.com/fwchen/jellyfish/domain/user/repository"
	"github.com/fwchen/jellyfish/domain/user/util"
	"github.com/fwchen/jellyfish/domain/visitor"
	"github.com/juju/errors"
)

func NewGuardImpl(userRepository repository.Repository) Guard {
	return &GuardImpl{userRepository: userRepository}
}

type GuardImpl struct {
	userRepository repository.Repository
}

func (g *GuardImpl) AuthenticateGuest(guest *visitor.Visitor) error {
	guest.IsCertified = true
	return g.logGuestAccess(guest)
}

func (g *GuardImpl) logGuestAccess(guest *visitor.Visitor) error {
	return nil
}

func (g *GuardImpl) GenerateUserHash(guest *visitor.Visitor) (string, error) {
	hash, err := util.GenerateFromPassword(guest.Password)
	if err != nil {
		return hash, errors.Trace(err)
	}
	return hash, nil
}

func (g *GuardImpl) compareGuestPassword(guest *visitor.Visitor) error {
	hash, err := g.userRepository.FindUserHashByName(guest.Name)
	if err != nil {
		return errors.Trace(err)
	}
	return util.CompareHashAndPassword(hash, guest.Password)
}
