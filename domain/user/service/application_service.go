package service

import (
	"jellyfish/domain/user/repository"
	"jellyfish/domain/user/response"
	"jellyfish/service"
	"github.com/juju/errors"
)

func NewApplicationService(userRepo repository.Repository, imageStorageService *service.ImageStorageService) *ApplicationService {
	return &ApplicationService{userRepo: userRepo, imageStorageService: imageStorageService}
}

type ApplicationService struct {
	userRepo            repository.Repository
	imageStorageService *service.ImageStorageService
}

func (a *ApplicationService) UpdateUserAvatar(userID string, avatar string) error {
	user, err := a.userRepo.FindByID(userID)
	if err != nil {
		return errors.Trace(err)
	}
	fileName, err := a.imageStorageService.SaveBase64Image(avatar)
	if err != nil {
		return errors.Trace(err)
	}
	user.SetAvatar(fileName)
	err = a.userRepo.Save(user)
	return err
}

func (a *ApplicationService) GetUserInfo(userID string) (*response.UserInfo, error) {
	user, err := a.userRepo.FindByID(userID)
	if err != nil {
		return nil, errors.Trace(err)
	}
	return response.TransformToUserInfo(user), nil
}