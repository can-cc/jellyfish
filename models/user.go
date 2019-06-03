package models

import (
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:createdAt`
}
