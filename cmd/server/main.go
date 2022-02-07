package main

import (
	"flag"
	"fmt"
	_ "github.com/labstack/gommon/log"
	"jellyfish/application"
	configs "jellyfish/config"
	"jellyfish/database"
	"jellyfish/logger"
	"jellyfish/service"
)

func main() {
	fmt.Println("Running server")
	configPath := flag.String("config", "config/config.yaml", "config file path")
	fmt.Printf("reading config file [%s]", *configPath)
	config, err := configs.LoadConfig(*configPath)
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
