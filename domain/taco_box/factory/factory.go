package factory

import (
	"github.com/fwchen/jellyfish/domain/taco_box"
)

func NewTacoBox(name, icon, creatorID string) *taco_box.TacoBox {
	return &taco_box.TacoBox{
		Name:      name,
		Icon:      icon,
		CreatorID: creatorID,
	}
}
