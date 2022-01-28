package service

import (
	"jellyfish/application/middleware"
	configs "jellyfish/config"
	"jellyfish/domain/visitor/factory"
	"jellyfish/domain/visitor/repository"
	"jellyfish/domain/visitor/service/impl"
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
	if !visitor.IsCertified {
		return nil, errors.New("username or password not match")
	}
	token, err := middleware.SignedToken(middleware.SignData{ID: id}, a.config.JwtSecret)
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
