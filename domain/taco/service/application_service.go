package service

import (
	"github.com/fwchen/jellyfish/domain/taco"
	"github.com/fwchen/jellyfish/domain/taco/command"
	"github.com/fwchen/jellyfish/domain/taco/factory"
	"github.com/fwchen/jellyfish/domain/taco/repository"
	"github.com/fwchen/jellyfish/domain/taco_box"
	"github.com/fwchen/jellyfish/domain/taco_box/service"
	"github.com/juju/errors"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func NewTacoApplicationService(tacoRepo repository.Repository, tacoBoxPermissionService *service.TacoBoxPermissionService) *TacoApplicationService {
	return &TacoApplicationService{
		tacoRepo:                 tacoRepo,
		tacoBoxPermissionService: tacoBoxPermissionService,
	}
}

type TacoApplicationService struct {
	tacoRepo                 repository.Repository
	tacoBoxPermissionService *service.TacoBoxPermissionService
}

func (t *TacoApplicationService) GetTacos(userID string, filter repository.ListTacoFilter) ([]taco.Taco, error) {
	return t.tacoRepo.ListTacos(userID, filter)
}

func (t *TacoApplicationService) CreateTaco(command *command.CreateTacoCommand, userID string) (*string, error) {
	if command.BoxId != nil {
		if !contains(taco_box.CommonTacoBoxes[:], *command.BoxId) {
			can, err := t.tacoBoxPermissionService.CheckUserCanOperation(*command.BoxId, userID)
			if err != nil {
				return nil, errors.Trace(err)
			}
			if !can {
				return nil, errors.Forbiddenf("user [userId = %s] forbidden create taco in box [boxId = %s]", userID, *command.BoxId)
			}
		}
	}
	ta := factory.NewTacoFromCreateCommand(command, userID)
	return t.tacoRepo.SaveTaco(ta)
}

func (t *TacoApplicationService) UpdateTaco(command command.UpdateTacoCommand) error {
	ta, err := t.tacoRepo.FindTaco(command.TacoID)
	if err != nil {
		return errors.Trace(err)
	}
	ta.Content = command.Content
	ta.Detail = command.Detail
	ta.Deadline = command.Deadline
	ta.Status = command.Status
	_, err = t.tacoRepo.SaveTaco(ta)
	return err
}
