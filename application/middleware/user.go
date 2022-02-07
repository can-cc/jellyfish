package middleware

import (
	"github.com/labstack/echo"
)

func GetUserID(c echo.Context) string {
	req := c.Request()
	headers := req.Header
	userId := headers.Get("x-app-user-id")
	return userId
}
