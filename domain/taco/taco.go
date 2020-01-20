package taco

import (
	"strings"
	"time"
)

type Status string

const (
	Doing Status = "Doing"
	Done  Status = "Done"
)

func ParseStatues(str string) (statues []Status) {
	if str == "" {
		return
	}
	ss := strings.Split(str, ",")
	for _, s := range ss {
		statues = append(statues, Status(strings.Trim(s, " ")))
	}
	return
}

type Type string

const (
	Task Type = "Task"
)

type Taco struct {
	ID        string     `json:"id"`
	CreatorID string     `json:"creatorID"`
	Content   string     `json:"content"`
	Detail    *string    `json:"detail"`
	Status    Status     `json:"status"`
	Type      Type       `json:"type"`
	Deadline  *time.Time `json:"deadline"`
	BoxId     *string    `json:"boxID"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdateAt  *time.Time `json:"updatedAt"`
}

func (t *Taco) IsNew() bool {
	return t.ID == ""
}
