package service

import (
	"github.com/fwchen/jellyfish/domain/taco"
	"github.com/fwchen/jellyfish/domain/taco/command"
	"github.com/fwchen/jellyfish/domain/taco/factory"
	"github.com/fwchen/jellyfish/domain/taco/repository"
	"github.com/fwchen/jellyfish/domain/taco_box"
	"github.com/fwchen/jellyfish/domain/taco_box/service"
	"github.com/fwchen/jellyfish/util"
	"github.com/juju/errors"
)

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

func (t *TacoApplicationService) GetTacos(userID string, status []taco.Status, box string) ([]taco.Taco, error) {
	var tacoTypeStr string
	var boxId *string = nil
	if taco_box.ContainTypeTacoBox(box) {
		tacoTypeStr = box
	} else if box == taco_box.TacoBoxAll {
		tacoTypeStr = string(taco.Task)
	} else {
		tacoTypeStr = string(taco.Task)
		if box == "" {
			boxId = nil
		} else {
			boxId = util.PointerStr(box)
		}
	}
	tacoType := taco.Type(tacoTypeStr)
	return t.tacoRepo.List(userID, taco.ListTacoFilter{
		Statues: status,
		Type:    &tacoType,
		BoxId:   boxId,
	})
}

func (t *TacoApplicationService) CreateTaco(command *command.CreateTacoCommand, userID string) (*string, error) {
	maxOrder, err := t.tacoRepo.FindMaxOrderByCreatorID(userID)
	if err != nil {
		return nil, errors.Trace(err)
	}
	command.Order = *maxOrder + float64(10)
	if command.BoxId != nil {
		if !taco_box.ContainCommonTacoBox(*command.BoxId) {
			can, err := t.tacoBoxPermissionService.CheckUserCanOperation(*command.BoxId, userID)
			if err != nil {
				return nil, errors.Trace(err)
			}
			if !can {
				return nil, errors.Forbiddenf("user [userId = %s] forbidden create taco in box [boxId = %s]", userID, *command.BoxId)
			}
		} else {
			command.BoxId = nil
		}
	}
	ta := factory.NewTacoFromCreateCommand(command, userID)
	return t.tacoRepo.Save(ta)
}

func (t *TacoApplicationService) UpdateTaco(command command.UpdateTacoCommand) error {
	ta, err := t.tacoRepo.FindById(command.TacoId)
	if err != nil {
		return errors.Trace(err)
	}
	ta.Content = command.Content
	ta.Detail = command.Detail
	ta.Deadline = command.Deadline
	ta.Status = command.Status
	ta.BoxId = command.BoxId
	_, err = t.tacoRepo.Save(ta)
	return err
}

func (t *TacoApplicationService) DeleteTaco(id string) error {
	return t.tacoRepo.Delete(id)
}
