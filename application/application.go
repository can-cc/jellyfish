package application

import (
	"fmt"
	configs "github.com/fwchen/jellyfish/config"
	"github.com/fwchen/jellyfish/database"
	"github.com/juju/errors"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewApplication(config *configs.AppConfig, datasource *database.AppDataSource) Application {
	return Application{
		config:     config,
		datasource: datasource,
	}
}

type Application struct {
	config     *configs.AppConfig
	datasource *database.AppDataSource
}

func (a *Application) StartServe() {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())

	a.Route(e)

	e.HTTPErrorHandler = func(err error, context echo.Context) {
		// TODO config.DeployEnv 来判断一下
		e.DefaultHTTPErrorHandler(err, context)
		fmt.Println()
		fmt.Println(errors.ErrorStack(err))
		fmt.Println()
	}

	fmt.Println(fmt.Sprintf("jellyfish serve on http://%s", a.config.Application.Addr))
	e.Logger.Fatal(e.Start(a.config.Application.Addr))
}
