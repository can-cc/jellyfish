package main

import (
	"jellyfish/application"
	configs "jellyfish/config"
	"jellyfish/database"
	"jellyfish/logger"
	"jellyfish/service"
	_ "github.com/labstack/gommon/log"
	"github.com/opentracing/opentracing-go"
	"go.elastic.co/apm/module/apmot"
)

func main() {
	config, err := configs.LoadConfig("config/config.yaml")
	if err != nil {
		panic(err)
	}

	opentracing.SetGlobalTracer(apmot.New())

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
