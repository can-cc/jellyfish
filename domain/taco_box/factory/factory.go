package factory

import (
	"jellyfish/domain/taco_box"
)

func NewTacoBox(name, creatorID string) *taco_box.TacoBox {
	return &taco_box.TacoBox{
		Name:      name,
		CreatorID: creatorID,
	}
}
