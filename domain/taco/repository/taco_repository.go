package repository

import "github.com/fwchen/jellyfish/domain/taco"

type Repository interface {
	List(userID string, filter taco.ListTacoFilter) ([]taco.Taco, error)
	Save(taco *taco.Taco) (*string, error)
	FindById(tacoID string) (*taco.Taco, error)
	Delete(tacoId string) error
}
