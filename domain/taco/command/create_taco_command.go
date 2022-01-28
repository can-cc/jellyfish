package command

import (
	"time"

	"jellyfish/domain/taco_box"
)

type CreateTacoCommand struct {
	Content  string     `json:"content" validate:"required"`
	Detail   *string    `json:"detail"`
	Deadline *time.Time `json:"deadline"`
	BoxId    *string    `json:"boxId"`
	Order    float64
}

type SortTacoCommand struct {
	TacoId       string              `json:"tacoId"`
	TargetTacoId string              `json:"targetTacoId"`
	IsBefore     bool                `json:"isBefore"`
	BoxId        *taco_box.TacoBoxID `json:"boxId"`
}
