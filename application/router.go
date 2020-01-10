package application

import (
	"github.com/dchest/captcha"
	appMiddleware "github.com/fwchen/jellyfish/application/middleware"
	userHandler "github.com/fwchen/jellyfish/domain/user/handler"
	userRepoImpl "github.com/fwchen/jellyfish/domain/user/repository/impl"
	visitorHandler "github.com/fwchen/jellyfish/domain/visitor/handler"
	"github.com/fwchen/jellyfish/domain/visitor/repository/impl"
	"net/http"

	"github.com/labstack/echo"
)

func (a *application) Route(e *echo.Echo) {
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	authorizeGroup := appMiddleware.ApplyJwtInRoute(e, &a.config.Application)

	{
		handler := visitorHandler.NewHandler(impl.NewVisitorRepository(a.datasource), &a.config.Application)
		e.POST("/login", handler.Login)
		e.POST("/signup", handler.SignUp)
		e.GET("/captcha/*", echo.WrapHandler(captcha.Server(captcha.StdWidth, captcha.StdHeight)))
		e.POST("/captcha", handler.GenCaptcha)
	}

	{
		handler := userHandler.NewHandler(userRepoImpl.NewUserRepository(a.datasource))
		authUserGroup := authorizeGroup.Group("user")
		authUserGroup.GET("/:userID", handler.GetUserInfo)
		authUserGroup.POST("/avatar", handler.UpdateUserAvatar)
	}
}
