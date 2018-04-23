package models

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"fmt"

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

	result, err2 := stmt.Exec(user.Username, hash, time.Now().Unix())

	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}

func CheckUserExist(db *sql.DB, username string) bool {
	sql := "SELECT hash FROM users WHERE username = ?"
	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	rows, err2 := stmt.Query(username)

	if err2 != nil {
		panic(err)
	}

	return rows.Next()
}

func CompareHashAndPassword(db *sql.DB, user *User) bool {
	sql := "SELECT hash FROM users WHERE username = ?"
	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	// Replace the '?' in our prepared statement with 'id'
	rows, err2 := stmt.Query(user.Username)

	if err2 != nil {
		panic(err2)
	}

	var hash string
	rows.Next()
	rows.Scan(hash)

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(user.Password))
	if err != nil {
		return false
	} else {
		return true
	}
}
