package application

import (
	"github.com/dchest/captcha"
	appMiddleware "github.com/fwchen/jellyfish/application/middleware"
	tacoHandler "github.com/fwchen/jellyfish/domain/taco/handler"
	tacoRepoImpl "github.com/fwchen/jellyfish/domain/taco/repository/impl"
	userHandler "github.com/fwchen/jellyfish/domain/user/handler"
	userRepoImpl "github.com/fwchen/jellyfish/domain/user/repository/impl"
	visitorHandler "github.com/fwchen/jellyfish/domain/visitor/handler"
	visitorRepoImpl "github.com/fwchen/jellyfish/domain/visitor/repository/impl"
	"net/http"

	"github.com/labstack/echo"
)

func (a *Application) Route(e *echo.Echo) {

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	authorizeGroup := appMiddleware.ApplyJwtInRoute(e, &a.config.Application)

	{
		handler := visitorHandler.NewHandler(visitorRepoImpl.NewVisitorRepository(a.datasource), &a.config.Application)
		e.POST("/login", handler.Login)
		e.POST("/signup", handler.SignUp)
		e.GET("/captcha/*", echo.WrapHandler(captcha.Server(captcha.StdWidth, captcha.StdHeight)))
		e.POST("/captcha", handler.GenCaptcha)
	}

	{
		handler := userHandler.NewHandler(userRepoImpl.NewUserRepository(a.datasource))
		authUserGroup := authorizeGroup.Group("user")
		authUserGroup.GET("/me", handler.GetUserInfo)
		authUserGroup.POST("/avatar", handler.UpdateUserAvatar)
		e.GET("/avatar/:userID", handler.GetUserAvatar)
	}

	{
		handler := tacoHandler.NewHandler(tacoRepoImpl.NewTacoRepository(a.datasource))
		tacoGroup := authorizeGroup.Group("taco")
		tacoGroup.GET("s", handler.GetTacos)
		tacoGroup.POST("", handler.CreateTaco)
	}
}
