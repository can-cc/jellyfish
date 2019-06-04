package userrepository

import (
	"database/sql"
	"github.com/fwchen/jellyfish/models"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// CreateUser :
func CreateUser(db *sql.DB, user *models.User) (int64, error) {
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

// CheckUserExist :
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

// GetUserWhenCompareHashAndPassword :
func GetUserWhenCompareHashAndPassword(db *sql.DB, username string, password string) (models.User, error) {
	sql := "SELECT id, hash, created_at FROM users WHERE username = ?"
	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	rows, err2 := stmt.Query(username)

	if err2 != nil {
		panic(err2)
	}

	defer rows.Close()

	user := models.User{}
	user.Username = username
	var hash string
	rows.Next()
	rows.Scan(&user.ID, &hash, &user.CreatedAt)

	err3 := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return user, err3
}
