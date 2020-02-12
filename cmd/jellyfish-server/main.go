package main

import (
	"github.com/fwchen/jellyfish/application"
	configs "github.com/fwchen/jellyfish/config"
	"github.com/fwchen/jellyfish/database"
	"github.com/fwchen/jellyfish/logger"
	"github.com/fwchen/jellyfish/service"
	_ "github.com/labstack/gommon/log"
)

func main() {
	config, err := configs.LoadConfig("config/config.yaml")
	if err != nil {
		panic(err)
	}

	err = logger.InitLogger(config.Logger)
	if err != nil {
		panic(err)
	}
	datasource, err := database.GetDatabase(config.DataSource)
	if err != nil {
		panic(err)
	}

	storageService := service.NewStorageService(&config.Storage)
	err = storageService.Init()
	if err != nil {
		panic(err)
	}
	imageStorageService, err := service.NewImageStorageService(config.Bucket.Image, storageService)
	if err != nil {
		panic(err)
	}

	app := application.NewApplication(config, datasource, imageStorageService)
	app.StartServe()
}
