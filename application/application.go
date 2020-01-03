package application

import (
	"fmt"
	configs "github.com/fwchen/jellyfish/config"
	"github.com/fwchen/jellyfish/database"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewApplication(config *configs.AppConfig, datasource *database.AppDataSource) application {
	return application{
		config:     config,
		datasource: datasource,
	}
}

type application struct {
	config     *configs.AppConfig
	datasource *database.AppDataSource
}

func (a *application) StartServe() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	Route(e)

	fmt.Println(fmt.Sprintf("jellyfish serve on http://%s", a.config.Application.Addr))
	e.Logger.Fatal(e.Start(a.config.Application.Addr))
}
