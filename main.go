package main

import (
	"bytes"
	"github.com/fwchen/jellyfish/database"
	"github.com/fwchen/jellyfish/router"

	"fmt"
	"io/ioutil"

	"github.com/labstack/echo/middleware"
	_ "github.com/labstack/gommon/log"
	"github.com/spf13/viper"

	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

func readConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	customConfigSrc, err := ioutil.ReadFile("config.custom.yaml")
	if err != nil {
		panic(err)
	}

	err2 := viper.ReadInConfig()

	if err2 != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err2))
	}

	err3 := viper.MergeConfig(bytes.NewBuffer(customConfigSrc))
	if err3 != nil {
		panic(err3)
	}
}

func main() {
	readConfig()

	database.CheckDBConnect()

	e := echo.New()

	e.Use(middleware.Logger())
	//e.Use(middleware.Recover())
	e.Static("/upload", "upload")

	router.Route(e)

	fmt.Println("jellyfish serve on http://0.0.0.0:8000")
	e.Logger.Fatal(e.Start("0.0.0.0:8000"))
}
