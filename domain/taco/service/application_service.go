package service

import (
	"github.com/fwchen/jellyfish/domain/taco"
	"github.com/fwchen/jellyfish/domain/taco/command"
	"github.com/fwchen/jellyfish/domain/taco/factory"
	"github.com/fwchen/jellyfish/domain/taco/repository"
	"github.com/juju/errors"
)

func NewTacoApplicationService(tacoRepo repository.Repository) *TacoApplicationService {
	return &TacoApplicationService{tacoRepo: tacoRepo}
}

type TacoApplicationService struct {
	tacoRepo repository.Repository
}

func (t *TacoApplicationService) GetTacos(userID string, filter repository.ListTacoFilter) ([]taco.Taco, error) {
	return t.tacoRepo.ListTacos(userID, filter)
}

func (t *TacoApplicationService) CreateTaco(command *command.CreateTacoCommand, userID string) (*string, error) {
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
