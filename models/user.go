package models

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CreateUser(db *sql.DB, user *User) (int64, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	sql := "INSERT INTO users(username, hash, created_at) VALUES (?, ?, ?)"
	stmt, err := db.Prepare(sql)
	defer stmt.Close()

	result, err2 := stmt.Exec(user.Password, hash, time.Now().Unix())

	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}
