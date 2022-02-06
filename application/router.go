package application

import (
	"jellyfish/application/route"
	tacoHandler "jellyfish/domain/taco/handler"
	tacoRepoImpl "jellyfish/domain/taco/repository/impl"
	tacoBoxHandler "jellyfish/domain/taco_box/handler"
	tacoBoxImpl "jellyfish/domain/taco_box/repository/impl"
	"jellyfish/domain/taco_box/service"
	"net/http"

	"github.com/labstack/echo"
)

func (a *Application) Route(e *echo.Echo) {

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	e.GET("/image/:fileName", route.GetImageRoute(a.imageStorageService))

	tacoBoxRepo := tacoBoxImpl.NewTacoBoxRepositoryImpl(a.datasource)
	tacoBoxPermissionService := service.NewTacoBoxPermissionService(tacoBoxRepo)
	tbHandler := tacoBoxHandler.NewHandler(tacoBoxRepo)
	e.GET("/box/es", tbHandler.GetTacoBoxes)
	e.POST("/box", tbHandler.CreateTacoBox)
	e.PUT("/box/:tacoBoxID", tbHandler.UpdateTacoBox)

	tHandler := tacoHandler.NewHandler(tacoRepoImpl.NewTacoRepository(a.datasource), tacoBoxPermissionService)
	e.GET("/tacos", tHandler.GetTacos)
	e.POST("/taco", tHandler.CreateTaco)
	e.POST("/taco/resort", tHandler.SortTaco)
	e.PUT("/taco/:tacoId", tHandler.UpdateTaco)
	e.DELETE("/taco/:tacoId", tHandler.DeleteTaco)
}
