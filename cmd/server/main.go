package main

import (
	"flag"
	"fmt"
	_ "github.com/labstack/gommon/log"
	"jellyfish/application"
	configs "jellyfish/config"
	"jellyfish/database"
	"jellyfish/logger"
)

func main() {
	fmt.Println("Running server")
	configPath := flag.String("config", "config/config.yaml", "config file path") //  -config=xxx.yaml
	flag.Parse()
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

	app := application.NewApplication(config, datasource)
	app.StartServe()
}
