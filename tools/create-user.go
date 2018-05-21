package main

import (
	"fmt"
	"jellyfish/database"
	"jellyfish/models"
	"os"
)

func main() {
	var username string = os.Args[1]
	var password string = os.Args[2]

	db := database.InitDB("storage.sqlite3")
	defer db.Close()

	user := models.User{}

	user.Username = username
	user.Password = password

	models.CreateUser(db, &user)

	fmt.Print("Create user successful:")
	fmt.Print("username: ", username, "\npassword: ", password, "\n")

}
