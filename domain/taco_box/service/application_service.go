package service

import (
	"github.com/fwchen/jellyfish/domain/taco_box"
	"github.com/fwchen/jellyfish/domain/taco_box/command"
	"github.com/fwchen/jellyfish/domain/taco_box/factory"
	"github.com/fwchen/jellyfish/domain/taco_box/repository"
)

func NewTacoBoxApplicationService(repo repository.TacoBoxRepository) *TacoBoxApplicationService {
	return &TacoBoxApplicationService{tacoBoxRepo: repo}
}

type TacoBoxApplicationService struct {
	tacoBoxRepo repository.TacoBoxRepository
}

func (t *TacoBoxApplicationService) CreateTacoBox(command command.CreateTacoBoxCommand) (*taco_box.TacoBoxID, error) {
	taco := factory.NewTacoBox(command.Name, command.Icon, command.CreatorId)
	return t.tacoBoxRepo.SaveTacoBox(taco)
}
