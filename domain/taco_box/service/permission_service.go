package service

import "github.com/fwchen/jellyfish/domain/taco_box/repository"

func NewTacoBoxPermissionService(repo repository.TacoBoxRepository) *TacoBoxPermissionService {
	return &TacoBoxPermissionService{tacoBoxRepo: repo}
}

type TacoBoxPermissionService struct {
	tacoBoxRepo repository.TacoBoxRepository
}

func (t *TacoBoxPermissionService) CheckUserCanOperation(tacoBoxId string, userId string) (bool, error) {
	box, err := t.tacoBoxRepo.FindTacoBox(tacoBoxId)
	if err != nil {
		return false, nil
	}
	return box.CreatorID == userId, nil
}
