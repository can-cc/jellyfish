package handlers

import (
	"database/sql"
	"time"

	"github.com/dgrijalva/jwt-go"
	"jellyfish/models"
	"net/http"

	"github.com/dchest/captcha"

	"fmt"

	"github.com/labstack/echo"
)

type H2 map[string]interface{}

func SignUp(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := models.User{}

		request := new(struct {
			Captcha   string `json:"captcha"`
			CaptchaId string `json:"captchaId"`
			Username  string `json:"username"`
			Password  string `json:"password"`
		})

		c.Bind(&request)
		user.Username = request.Username
		user.Password = request.Password

		if !captcha.VerifyString(request.CaptchaId, request.Captcha) {
			fmt.Println(request)
			return c.NoContent(http.StatusUnauthorized)
		} else {
			_, error := models.CreateUser(db, &user)
			if error == nil {
				return c.NoContent(http.StatusNoContent)
			} else {
				return error
			}
		}

	}
}

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
		claims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()

		// TODO replace secret

		t, err := token.SignedString([]byte("secret"))

		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
			"id":    user.ID,
		})
	}
}
