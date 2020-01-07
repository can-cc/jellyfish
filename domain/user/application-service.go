package user

import (
	"github.com/fwchen/jellyfish/domain/user/repository"
	"github.com/fwchen/jellyfish/domain/user/response"
	"github.com/juju/errors"
)

func NewApplicationService(userRepo repository.Repository) *ApplicationService {
	return &ApplicationService{userRepo: userRepo}
}

type ApplicationService struct {
	userRepo repository.Repository
}

func (a *ApplicationService) UpdateUserAvatar(userID string, avatar string) error {
	user, err := a.userRepo.FindByID(userID)
	if err != nil {
		return errors.Trace(err)
	}
	user.SetAvatar(avatar)
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
