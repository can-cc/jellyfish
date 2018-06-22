package handlers

import (
	"database/sql"

	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func PostAvatar(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["id"].(string)

		form, err := c.MultipartForm()

		if err != nil {
			return err
		}
		files := form.File["avatar"]
		fmt.Println(len(files))

		if len(files) != 1 {
			return c.NoContent(http.StatusBadRequest)
		}

		for _, file := range files {
			// Source
			src, err := file.Open()
			if err != nil {
				return err
			}
			defer src.Close()

			// Destination
			dst, err := os.Create(file.Filename)
			if err != nil {
				return err
			}
			defer dst.Close()

			// Copy
			if _, err = io.Copy(dst, src); err != nil {
				return err
			}

		}
		return c.JSON(http.StatusUnauthorized, "")
	}
}
