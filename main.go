package main

import (
	"bytes"
	"net/http"

	"github.com/fwchen/jellyfish/database"
	"github.com/fwchen/jellyfish/handlers"

	"fmt"
	"io/ioutil"

	"github.com/dchest/captcha"

	_ "github.com/labstack/gommon/log"
	"github.com/spf13/viper"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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

	db := database.InitDB("storage.sqlite3?parseTime=true&cache=shared&mode=rwc")
	defer db.Close()

	database.Migrate(db)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/upload", "upload")

	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello my friends")
	})
	e.POST("/signin", handlers.SignIn(db))
	e.POST("/signup", handlers.SignUp(db))
	e.GET("/captcha/*", echo.WrapHandler(captcha.Server(captcha.StdWidth, captcha.StdHeight)))
	e.POST("/captcha", handlers.GenCaptcha(db))

	r := e.Group("/auth")
	r.Use(middleware.JWT(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte(viper.GetString("jwt-key")),
		TokenLookup: "header:App-Authorization",
	})))

	r.GET("/todo", handlers.GetUserTodos(db))
	r.GET("/user/:userId", handlers.GetUserInfo(db))
	r.POST("/todo", handlers.CreateTodo(db))
	r.DELETE("/todo/:id", handlers.DeleteTodo(db))
	r.PUT("/todo/:id", handlers.UpdateTodo(db))

	r.POST("/avatar", handlers.PostAvatar(db))
	r.POST("/avatar/base64", handlers.PostAvatarByBase64(db))

	fmt.Println("jellyfish serve on http://0.0.0.0:8000")
	e.Logger.Fatal(e.Start("0.0.0.0:8000"))
}
