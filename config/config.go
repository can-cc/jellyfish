package config

import (
	"database/sql"
	"time"

	"github.com/dgrijalva/jwt-go"
	"jellyfish/models"
	"net/http"

	"github.com/dchest/captcha"

	"encoding/json"
	"fmt"
	"os"

	"github.com/labstack/echo"
)

type Configuration struct {
	JwtSercetToken string
	Groups         []string
}

func readConfig() {

	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(configuration.Users) // output: [UserA, UserB]

}
