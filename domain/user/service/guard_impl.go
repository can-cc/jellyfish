package service

import (
	"github.com/fwchen/jellyfish/domain/user"
	"github.com/fwchen/jellyfish/domain/user/repository"
	"github.com/fwchen/jellyfish/domain/user/util"
	"github.com/juju/errors"
)

func NewGuardImpl(userRepository repository.UserRepository) Guard {
	return &GuardImpl{userRepository: userRepository}
}

type GuardImpl struct {
	userRepository repository.UserRepository
}

func (g *GuardImpl) AuthenticateGuest(guest *user.Guest) error {
	guest.IsCertified = true
	return g.logGuestAccess(guest)
}

func (g *GuardImpl) logGuestAccess(guest *user.Guest) error {
	return nil
}

func (g *GuardImpl) GenerateUserHash(guest *user.Guest) (string, error) {
	hash, err := util.GenerateFromPassword(guest.Password)
	if err != nil {
		return hash, errors.Trace(err)
	}
	return hash, nil
}

func (g *GuardImpl) compareGuestPassword(guest *user.Guest) error {
	hash, err := g.userRepository.FindUserHash(guest.Name)
	if err != nil {
		return errors.Trace(err)
	}
	return util.CompareHashAndPassword(hash, guest.Password)
}
