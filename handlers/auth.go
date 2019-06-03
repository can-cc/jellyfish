package handlers

import (
	"database/sql"
	"github.com/fwchen/jellyfish/repository/user"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/fwchen/jellyfish/models"
	"net/http"

	"github.com/dchest/captcha"

	"fmt"

	"github.com/labstack/echo"
)

type JwtCustomClaims struct {
	Username string `json:"username"`
	ID       string `json:"id"`
	jwt.StandardClaims
}

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
			_, error := user_repository.CreateUser(db, &user)
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

		isExist := user_repository.CheckUserExist(db, request.Username)

		if !isExist {
			return c.JSON(http.StatusBadRequest, "")
		}

		user, err := user_repository.GetUserWhenCompareHashAndPassword(db, request.Username, request.Password)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, "")
		}

		claims := &JwtCustomClaims{
			user.Username,
			user.ID,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
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
