package service

import (
	configs "github.com/fwchen/jellyfish/config"
	"github.com/fwchen/jellyfish/domain/visitor/factory"
	"github.com/fwchen/jellyfish/domain/visitor/repository"
	"github.com/fwchen/jellyfish/util"
	"github.com/juju/errors"
)

func NewApplicationService(visitorRepo repository.Repository, config *configs.ApplicationConfig) *ApplicationService {
	return &ApplicationService{visitorRepo: visitorRepo, config: config}
}

type ApplicationService struct {
	visitorRepo repository.Repository
	guard       Guard
	config      *configs.ApplicationConfig
}

func (a *ApplicationService) Login(username, password string) bool {
	visitor := factory.NewVisitor(username, password)
	a.guard.Authenticate(visitor, "")
	return visitor.IsCertified
}

func (a *ApplicationService) SignUp(username, password string) (*string, error) {
	visitor := factory.NewVisitor(username, password)
	visitor.IsCertified = true
	hash, err := a.guard.GeneratePasswordHash(visitor)
	if err != nil {
		return nil, errors.Trace(err)
	}
	id, err := a.visitorRepo.Save(visitor, hash)
	if err != nil {
		return nil, errors.Trace(err)
	}
	token, err := util.SignedToken(util.SignData{ID: id})
	if err != nil {
		return nil, errors.Trace(err)
	}
	return &token, nil
}
