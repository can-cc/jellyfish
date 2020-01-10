package command

import (
	"time"
)

type CreateTacoCommand struct {
	Content  string     `json:"content"`
	Detail   string     `json:"detail"`
	Deadline *time.Time `json:"deadline"`
}
