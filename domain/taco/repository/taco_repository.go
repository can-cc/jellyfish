package repository

import "jellyfish/domain/taco"

type Repository interface {
	List(userId string, filter taco.TacoFilter) ([]taco.Taco, error)
	Save(taco *taco.Taco) (*string, error)
	SaveList(tacos []taco.Taco) error
	FindById(tacoID string) (*taco.Taco, error)
	MaxOrderByCreatorId(userId string) (*float64, error)
	MaxOrderByBoxId(userId string) (*float64, error)
	Delete(tacoId string) error
}
