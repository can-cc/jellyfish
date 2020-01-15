package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	configs "github.com/fwchen/jellyfish/config"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"time"
)

type SignData struct {
	ID string `json:"id"`
}

type JwtAppClaims struct {
	SignData
	jwt.StandardClaims
}

func ApplyJwtInRoute(e *echo.Echo, config *configs.ApplicationConfig) *echo.Group {
	authorizeGroup := e.Group("")
	authorizeGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:      &JwtAppClaims{},
		ContextKey:  "user",
		SigningKey:  []byte(config.JwtSecret),
		TokenLookup: fmt.Sprintf("header:%s", config.JwtHeaderKey),
	}))
	return authorizeGroup
}

func GetClaimsUserID(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtAppClaims)
	return claims.ID
}

func SignedToken(data SignData, jwtSecretKey string) (string, error) {
	claims := &JwtAppClaims{
		data,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(), // TODO: configure
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecretKey))
}
