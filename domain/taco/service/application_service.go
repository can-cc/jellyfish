package service

import (
	"github.com/fwchen/jellyfish/domain/taco"
	"github.com/fwchen/jellyfish/domain/taco/repository"
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
