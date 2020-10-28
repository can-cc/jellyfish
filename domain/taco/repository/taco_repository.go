package repository

import "github.com/fwchen/jellyfish/domain/taco"

type Repository interface {
	List(userID string, filter taco.ListTacoFilter) ([]taco.Taco, error)
	Save(taco *taco.Taco) (*string, error)
	SaveList(tacos []taco.Taco) error
	FindById(tacoID string) (*taco.Taco, error)
	MaxOrderByCreatorId(userId string) (*float64, error)
	MaxOrderByBoxId(userId string) (*float64, error)
	Delete(tacoId string) error
}
