package userrepository

import (
	"github.com/fwchen/jellyfish/database"
	"github.com/fwchen/jellyfish/models"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser :
func CreateUser(user *models.User) (string, error) {
	db := database.GetDB()

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	var id string
	sqlStatement := `
		INSERT INTO users (username, hash, created_at)
		VALUES ($1, $2, now()) RETURNING id`
	err2 := db.QueryRow(sqlStatement, user.Username, hash).Scan(&id)

	return id, err2
}

// CheckUserExist :
func CheckUserExist(username string) bool {
	db := database.GetDB()
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
func GetUserWhenCompareHashAndPassword(username string, password string) (models.User, error) {
	db := database.GetDB()
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
