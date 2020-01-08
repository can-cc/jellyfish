package handlers

import (
	user2 "github.com/fwchen/jellyfish/domain/user"
	"github.com/fwchen/jellyfish/repository"
	"time"

	"github.com/spf13/viper"

	"fmt"
	"net/http"

	"github.com/dchest/captcha"
	"github.com/dgrijalva/jwt-go"

	"github.com/labstack/echo"
)

func SignUp() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := user2.AppUser{}

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

			return c.String(http.StatusBadRequest, "captcha invalid")
		} else {
			_, error := repository.CreateUser(&user)
			if error == nil {
				return c.NoContent(http.StatusNoContent)
			} else {
				fmt.Print(error)
				return error
			}
		}

	}
}

func SignIn() echo.HandlerFunc {
	return func(c echo.Context) error {

		request := new(struct {
			Username string `json:"username"`
			Password string `json:"password"`
		})

		c.Bind(&request)

		isExist := repository.CheckUserExist(request.Username)

		if !isExist {
			return c.JSON(http.StatusBadRequest, "")
		}

		user, err := repository.GetUserWhenCompareHashAndPassword(request.Username, request.Password)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, "")
		}

		claims := &JwtAppClaims{
			user.Username,
			user.ID,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate 	encoded token and send it as response.
		t, err := token.SignedString([]byte(viper.GetString("JWT_SECRET")))
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
			"id":    user.ID,
		})
	}
}
