package taco

import (
	"time"
)

type Status string

const (
	Doing Status = "Doing"
	Done  Status = "Done"
)

type Type string

const (
	Task Type = "Task"
)

type Taco struct {
	ID        string     `json:"id"`
	CreatorID string     `json:"creatorID"`
	Content   string     `json:"content"`
	Detail    string     `json:"detail"`
	Status    Status     `json:"status"`
	Type      Type       `json:"type"`
	Deadline  *time.Time `json:"deadline"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdateAt  *time.Time `json:"updatedAt"`
}
