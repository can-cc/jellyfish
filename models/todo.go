package models

import (
	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	ID        int    `json:"id"`
	Content   string `json:"content"`
	Detail    string `json:"detail"`
	Deadline  int64  `json:"deadline"`
	Done      bool   `json:"done"`
	Status    string `json:"status"`
	Type      string `json:"type"`
	CreatorID string `json:"creatorId"`
	CreatedAt int64  `json:"createdAt"`
}

// TodoCollection :
type TodoCollection struct {
	Items []Todo `json:"items"`
}
