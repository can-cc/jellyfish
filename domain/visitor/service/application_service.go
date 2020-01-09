package service

import (
	"github.com/fwchen/jellyfish/application/middleware"
	configs "github.com/fwchen/jellyfish/config"
	"github.com/fwchen/jellyfish/domain/visitor/factory"
	"github.com/fwchen/jellyfish/domain/visitor/repository"
	"github.com/fwchen/jellyfish/domain/visitor/service/impl"
	"github.com/juju/errors"
)

func NewApplicationService(visitorRepo repository.Repository, config *configs.ApplicationConfig) *ApplicationService {
	return &ApplicationService{visitorRepo: visitorRepo, guard: &impl.GuardImpl{}, config: config}
}

type ApplicationService struct {
	visitorRepo repository.Repository
	guard       Guard
	config      *configs.ApplicationConfig
}

func (a *ApplicationService) Login(username, password string) (*string, error) {
	visitor := factory.NewVisitor(username, password)
	id, hash, err := a.visitorRepo.FindUserIDAndHash(visitor.Name)
	if err != nil {
		return nil, errors.Trace(err)
	}
	a.guard.Authenticate(visitor, hash)
	token, err := middleware.SignedToken(middleware.SignData{ID: id})
	if err != nil {
		return nil, errors.Trace(err)
	}
	return &token, nil
}

func (a *ApplicationService) SignUp(username, password string) error {
	visitor := factory.NewVisitor(username, password)
	visitor.IsCertified = true
	hash, err := a.guard.GeneratePasswordHash(visitor)
	if err != nil {
		return errors.Trace(err)
	}
	_, err = a.visitorRepo.Save(visitor, hash)
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}
