package application

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"time"
)

type SignData struct {
	ID string `json:"id"`
}

type JwtAppClaims struct {
	SignData
	jwt.StandardClaims
}

func GetClaimsUserID(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtAppClaims)
	return claims.ID
}

func SignedToken(data SignData) (string, error) {
	claims := &JwtAppClaims{
		data,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(viper.GetString("JWT_SECRET")))
}
