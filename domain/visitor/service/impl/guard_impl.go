package impl

import (
	"github.com/fwchen/jellyfish/domain/visitor"
	"github.com/fwchen/jellyfish/domain/visitor/util"
	"github.com/fwchen/jellyfish/logger"
	"github.com/juju/errors"
)

type GuardImpl struct {
}

func (g *GuardImpl) Authenticate(visitor *visitor.Visitor, hash string) {
	if g.compareGuestPassword(visitor, hash) == nil {
		visitor.IsCertified = true
		logger.L.Infow("visitor logged in success", "username=", visitor.Name)
	} else {
		visitor.IsCertified = false
		logger.L.Warnw("visitor logged in fail", "username=", visitor.Name)
	}
}

func (g *GuardImpl) GeneratePasswordHash(guest *visitor.Visitor) (string, error) {
	hash, err := util.GenerateFromPassword(guest.Password)
	if err != nil {
		return hash, errors.Trace(err)
	}
	return hash, nil
}

func (g *GuardImpl) compareGuestPassword(visitor *visitor.Visitor, hash string) error {
	return util.CompareHashAndPassword(hash, visitor.Password)
}
