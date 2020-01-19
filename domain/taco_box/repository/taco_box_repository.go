package repository

import "github.com/fwchen/jellyfish/domain/taco_box"

type TacoBoxRepository interface {
	SaveTacoBox(box *taco_box.TacoBox) (*taco_box.TacoBoxID, error)
	ListTacoBoxes(userID string) ([]taco_box.TacoBox, error)
	FindTacoBox(boxID string) (*taco_box.TacoBox, error)
}
