package userrepository

import (
	"github.com/fwchen/jellyfish/database"
	. "github.com/fwchen/jellyfish/models"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser :
func CreateUser(user *User) (string, error) {
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

	var hash string
	err := db.QueryRow(`SELECT  hash FROM users WHERE  username = $1`, username).Scan(&hash)

	return err == nil
}

// GetUserWhenCompareHashAndPassword :
func GetUserWhenCompareHashAndPassword(username string, password string) (*User, error) {
	db := database.GetDB()

	user := new(User)
	var hash string

	err := db.QueryRow(`SELECT id, hash, created_at FROM users WHERE username = $1`, username).Scan(&user.ID, &hash, &user.CreatedAt)
	if err != nil {
		panic(err)
	}

	err2 := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return user, err2
}
