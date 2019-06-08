package models

import (
	"time"
	"gopkg.in/guregu/null.v3"
)

type Todo struct {
	ID        string    `json:"id"`
	Content   string `json:"content"`
	Detail    string `json:"detail"`
	Deadline  null.Time  `json:"deadline",omitempty`
	Done      bool   `json:"done"`
	Status    string `json:"status"`
	Type      string `json:"type"`
	CreatorID string `json:"creatorId"`
	CreatedAt time.Time  `json:"createdAt"`
	updatedAt time.Time  `json:"updatedAt"`
}

// TodoCollection :
type TodoCollection struct {
	Items []Todo `json:"items"`
}
