package application

import (
	"fmt"
	configs "github.com/fwchen/jellyfish/config"
	"github.com/fwchen/jellyfish/database"
	"github.com/fwchen/jellyfish/service"
	"github.com/juju/errors"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.elastic.co/apm/module/apmecho"
	_ "go.elastic.co/apm/module/apmsql/sqlite3"
	"net/http"
	"runtime"
)

func NewApplication(config *configs.AppConfig, datasource *database.AppDataSource, imageStorageService *service.ImageStorageService) Application {
	return Application{
		config:              config,
		datasource:          datasource,
		imageStorageService: imageStorageService,
	}
}

type Application struct {
	config              *configs.AppConfig
	datasource          *database.AppDataSource
	imageStorageService *service.ImageStorageService
}

func (a *Application) StartServe() {
	e := echo.New()
	e.Use(apmecho.Middleware())

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())

	a.Route(e)

	e.HTTPErrorHandler = func(err error, context echo.Context) {
		stack := make([]byte, 4 << 10) // 4kb
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
