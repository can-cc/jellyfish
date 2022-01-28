package command

import (
	"jellyfish/domain/taco"
	"time"
)

type UpdateTacoCommand struct {
	TacoId          string      `json:"id" validate:"required"`
	Content         string      `json:"content" validate:"required"`
	Detail          *string     `json:"detail"`
	Deadline        *time.Time  `json:"deadline"`
	Status          taco.Status `json:"status"`
	BoxId           *string     `json:"boxId"`
	OperationUserID string
}
