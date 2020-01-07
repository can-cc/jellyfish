package application

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type JwtAppClaims struct {
	Username string `json:"username"`
	ID       string `json:"id"`
	jwt.StandardClaims
}

func GetClaimsUserID(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtAppClaims)
	return claims.ID
}
