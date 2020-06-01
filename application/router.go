package application

import (
	"github.com/dchest/captcha"
	appMiddleware "github.com/fwchen/jellyfish/application/middleware"
	"github.com/fwchen/jellyfish/application/route"
	tacoHandler "github.com/fwchen/jellyfish/domain/taco/handler"
	tacoRepoImpl "github.com/fwchen/jellyfish/domain/taco/repository/impl"
	tacoBoxHandler "github.com/fwchen/jellyfish/domain/taco_box/handler"
	tacoBoxImpl "github.com/fwchen/jellyfish/domain/taco_box/repository/impl"
	"github.com/fwchen/jellyfish/domain/taco_box/service"
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
		e.GET("/image/:fileName", route.GetImageRoute(a.imageStorageService))
	}

	{
		handler := visitorHandler.NewHandler(visitorRepoImpl.NewVisitorRepository(a.datasource), &a.config.Application)
		e.POST("/login", handler.Login)
		e.POST("/signup", handler.SignUp)
		e.GET("/captcha/*", echo.WrapHandler(captcha.Server(captcha.StdWidth, captcha.StdHeight)))
		e.POST("/captcha", handler.GenCaptcha)
	}

	{
		handler := userHandler.NewHandler(userRepoImpl.NewUserRepository(a.datasource), a.imageStorageService)
		authUserGroup := authorizeGroup.Group("user")
		authUserGroup.GET("/me", handler.GetUserInfo)
		authUserGroup.POST("/avatar", handler.UpdateUserAvatar)
	}

	tacoBoxRepo := tacoBoxImpl.NewTacoBoxRepositoryImpl(a.datasource)
	tacoBoxPermissionService := service.NewTacoBoxPermissionService(tacoBoxRepo)
	{
		handler := tacoBoxHandler.NewHandler(tacoBoxRepo)
		tacoBoxGroup := authorizeGroup.Group("box")
		tacoBoxGroup.GET("es", handler.GetTacoBoxes)
		tacoBoxGroup.POST("", handler.CreateTacoBox)
		tacoBoxGroup.PUT("/:tacoBoxID", handler.UpdateTacoBox)
	}

	{
		handler := tacoHandler.NewHandler(tacoRepoImpl.NewTacoRepository(a.datasource), tacoBoxPermissionService)
		tacoGroup := authorizeGroup.Group("taco")
		tacoGroup.GET("s", handler.GetTacos)
		tacoGroup.POST("", handler.CreateTaco)
		tacoGroup.PUT("/:tacoID", handler.UpdateTaco)
	}

}
