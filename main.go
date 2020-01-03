package main

import (
	"github.com/fwchen/jellyfish/application"
	configs "github.com/fwchen/jellyfish/config"
	"github.com/fwchen/jellyfish/database"
	"github.com/fwchen/jellyfish/logger"
	_ "github.com/labstack/gommon/log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	config, err := configs.LoadConfig("config/config.yaml")
	if err != nil {
		panic(err)
	}

	logger.InitLogger(config.Logger)
	datasource, err := database.GetDatabase(config.DataSource)
	if err != nil {
		panic(err)
	}

	app := application.NewApplication(config, datasource)
	app.StartServe()
}
