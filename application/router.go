package application

import (
	"github.com/dchest/captcha"
	userHandler "github.com/fwchen/jellyfish/domain/user/handler"
	userRepoImpl "github.com/fwchen/jellyfish/domain/user/repository/impl"
	"github.com/fwchen/jellyfish/handlers"
	"net/http"

	"github.com/spf13/viper"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func (a *application) Route(e *echo.Echo) {
	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello my friends")
	})
	e.POST("/signin", handlers.SignIn())
	e.POST("/signup", handlers.SignUp())
	e.GET("/captcha/*", echo.WrapHandler(captcha.Server(captcha.StdWidth, captcha.StdHeight)))
	e.POST("/captcha", handlers.GenCaptcha())

	authorizeGroup := e.Group("")
	authorizeGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:      &JwtAppClaims{},
		SigningKey:  []byte(viper.GetString("JWT_SECRET")),
		TokenLookup: "header:App-Authorization",
	}))

	{
		handler := userHandler.NewHandler(userRepoImpl.NewUserRepository(a.datasource))
		authUserGroup := authorizeGroup.Group("user")
		authUserGroup.GET("/:userID", handler.GetUserInfo)
	}

	authorizeGroup.GET("/todos", handlers.GetUserTodos())
	authorizeGroup.GET("/todos/done", handlers.GetUserDoneTodos())
	authorizeGroup.GET("/todos/doing", handlers.GetUserDoingTodos())

	authorizeGroup.POST("/todo", handlers.CreateTodo())
	authorizeGroup.DELETE("/todo/:id", handlers.DeleteTodo())
	authorizeGroup.PUT("/todo/:id", handlers.UpdateTodo())

	authorizeGroup.POST("/avatar", handlers.PostAvatar())
	authorizeGroup.POST("/avatar/base64", handlers.PostAvatarByBase64())
}
