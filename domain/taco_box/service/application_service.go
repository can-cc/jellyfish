package service

import (
	"github.com/fwchen/jellyfish/domain/taco_box"
	"github.com/fwchen/jellyfish/domain/taco_box/command"
	"github.com/fwchen/jellyfish/domain/taco_box/factory"
	"github.com/fwchen/jellyfish/domain/taco_box/repository"
	"github.com/juju/errors"
)

func NewTacoBoxApplicationService(repo repository.TacoBoxRepository) *TacoBoxApplicationService {
	return &TacoBoxApplicationService{tacoBoxRepo: repo}
}

type TacoBoxApplicationService struct {
	tacoBoxRepo repository.TacoBoxRepository
}

func (t *TacoBoxApplicationService) CreateTacoBox(command command.CreateTacoBoxCommand) (*taco_box.TacoBoxID, error) {
	taco := factory.NewTacoBox(command.Name, command.CreatorId)
	return t.tacoBoxRepo.SaveTacoBox(taco)
}

func (t *TacoBoxApplicationService) GetTacoBoxes(userID string) ([]taco_box.TacoBox, error) {
	return t.tacoBoxRepo.ListTacoBoxes(userID)
}

func (t *TacoBoxApplicationService) UpdateTacoBox(command command.UpdateTacoCommand) error {
	tb, err := t.tacoBoxRepo.FindTacoBox(command.TacoBoxID)
	if err != nil {
		return errors.Trace(err)
	}
	tb.Name = command.Name
	_, err = t.tacoBoxRepo.SaveTacoBox(tb)
	return err
}
