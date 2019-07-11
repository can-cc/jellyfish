package main

import (
	"github.com/fwchen/jellyfish/database"
	"github.com/fwchen/jellyfish/router"

	"fmt"

	"github.com/labstack/echo/middleware"
	_ "github.com/labstack/gommon/log"
	"github.com/spf13/viper"

	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

func readConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("JFISH")

	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s \n", err))
	}
}

func main() {
	readConfig()

	database.CheckDBConnect()

	e := echo.New()

	e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	e.Static("/upload", "upload")

	router.Route(e)

	fmt.Println("jellyfish serve on http://0.0.0.0:8180")
	e.Logger.Fatal(e.Start("0.0.0.0:8180"))
}
