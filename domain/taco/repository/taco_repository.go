package repository

import "github.com/fwchen/jellyfish/domain/taco"

type Repository interface {
	ListTacos(userID string, filter taco.ListTacoFilter) ([]taco.Taco, error)
	SaveTaco(taco *taco.Taco) (*string, error)
	FindTaco(tacoID string) (*taco.Taco, error)
}
