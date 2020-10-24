package repository

import "github.com/fwchen/jellyfish/domain/taco"

type Repository interface {
	List(userID string, filter taco.ListTacoFilter) ([]taco.Taco, error)
	Save(taco *taco.Taco) (*string, error)
	FindById(tacoID string) (*taco.Taco, error)
	//FindMaxOrderByBoxId() (float64, error)
	FindMaxOrderByCreatorID(userID string) (*float64, error)
	Delete(tacoId string) error
}
