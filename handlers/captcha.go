package handlers

import (
	"database/sql"

	"net/http"

	"github.com/dchest/captcha"

	"github.com/labstack/echo"
)

func GenCaptcha(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		d := struct {
			CaptchaId string
		}{
			captcha.New(),
		}

		return c.JSON(http.StatusCreated, map[string]string{
			"id": d.CaptchaId,
		})
	}
}
