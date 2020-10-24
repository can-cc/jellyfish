package command

import (
	"time"
)

type CreateTacoCommand struct {
	Content  string     `json:"content" validate:"required"`
	Detail   *string    `json:"detail"`
	Deadline *time.Time `json:"deadline"`
	BoxId    *string    `json:"boxId"`
	Order    float64
}
