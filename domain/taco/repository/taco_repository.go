package repository

import "github.com/fwchen/jellyfish/domain/taco"

type ListTacoFilter struct {
	Statues []taco.Status
}

type Repository interface {
	ListTacos(userID string, filter ListTacoFilter) ([]taco.Taco, error)
	SaveTaco(taco *taco.Taco) (*string, error)
	FindTaco(tacoID string) (*taco.Taco, error)
}
