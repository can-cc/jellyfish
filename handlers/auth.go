package handlers

import (
	"database/sql"
	"time"

	"fmt"
	"github.com/dgrijalva/jwt-go"
	"jellyfish/models"
	"net/http"

	"github.com/labstack/echo"
)

type H2 map[string]interface{}

func SignIn(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		request := new(struct {
			Username string `json:"username"`
			Password string `json:"password"`
		})

		c.Bind(&request)

		isExist := models.CheckUserExist(db, request.Username)

		if !isExist {
			return c.JSON(http.StatusBadRequest, "")
		}

		user, err := models.GetUserWhenCompareHashAndPassword(db, request.Username, request.Password)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, "")
		}

		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["username"] = user.Username
		claims["id"] = user.ID
		claims["createdAt"] = user.CreatedAt
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
		// TODO replace secret
		t, err := token.SignedString([]byte("secret"))

		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}
}
