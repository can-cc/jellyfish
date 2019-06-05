package router

import (
	"github.com/fwchen/jellyfish/handlers"
	"net/http"


	"github.com/dchest/captcha"

	"github.com/spf13/viper"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Route(e *echo.Echo) {


	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello my friends")
	})
	e.POST("/signin", handlers.SignIn())
	e.POST("/signup", handlers.SignUp())
	e.GET("/captcha/*", echo.WrapHandler(captcha.Server(captcha.StdWidth, captcha.StdHeight)))
	e.POST("/captcha", handlers.GenCaptcha())

	r := e.Group("/auth")
	r.Use(middleware.JWT(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte(viper.GetString("jwt-key")),
		TokenLookup: "header:App-Authorization",
	})))

	r.GET("/todo", handlers.GetUserTodos())
	r.GET("/user/:userId", handlers.GetUserInfo())
	r.POST("/todo", handlers.CreateTodo())
	r.DELETE("/todo/:id", handlers.DeleteTodo())
	r.PUT("/todo/:id", handlers.UpdateTodo())

	r.POST("/avatar", handlers.PostAvatar())
	r.POST("/avatar/base64", handlers.PostAvatarByBase64())
}