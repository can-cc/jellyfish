package command

import (
	"github.com/fwchen/jellyfish/domain/taco"
	"time"
)

type UpdateTacoCommand struct {
	Content         string      `json:"content" validate:"required"`
	Detail          *string     `json:"detail"`
	Deadline        *time.Time  `json:"deadline"`
	Status          taco.Status `json:"status"`
	TacoID          string
	OperationUserID string
}
