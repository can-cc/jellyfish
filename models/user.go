package models

import (
	"database/sql"
	// "fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt int64
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
	defer stmt.Close()

	rows, err2 := stmt.Query(username)

	if err2 != nil {
		panic(err)
	}
	defer rows.Close()

	return rows.Next()
}

func GetUserWhenCompareHashAndPassword(db *sql.DB, username string, password string) (User, error) {
	sql := "SELECT id, hash, created_at FROM users WHERE username = ?"
	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	// Replace the '?' in our prepared statement with 'id'
	rows, err2 := stmt.Query(username)

	if err2 != nil {
		panic(err2)
	}

	defer rows.Close()

	user := User{}
	user.Username = username
	var hash string
	rows.Next()
	rows.Scan(&user.ID, &hash, &user.CreatedAt)

	err3 := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return user, err3
}
