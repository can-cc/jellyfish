package application

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/juju/errors"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	configs "jellyfish/config"
	"jellyfish/database"
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
		stack := make([]byte, 4<<10) // 4kb
		length := runtime.Stack(stack, true)
		fmt.Printf("[PANIC RECOVER] %v %s\n", err, stack[:length])
		fmt.Println()
		if errors.IsBadRequest(err) {
			_ = context.NoContent(http.StatusBadRequest)
			return
		}
		if errors.IsForbidden(err) {
			_ = context.NoContent(http.StatusForbidden)
			return
		}
		e.DefaultHTTPErrorHandler(err, context)
	}

	fmt.Println(fmt.Sprintf("jellyfish serve on http://%s", a.config.Application.Addr))
	e.Logger.Fatal(e.Start(a.config.Application.Addr))
}
