package main

import (
	"database/sql"

	"bytes"
	"jellyfish/database"
	"jellyfish/handlers"
	"net/http"

	"fmt"
	"github.com/dchest/captcha"
	"io/ioutil"

	_ "github.com/labstack/gommon/log"
	"github.com/spf13/viper"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func migrate(db *sql.DB) {
	sql := `
    CREATE TABLE IF NOT EXISTS todos(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        creater_id INTEGER NOT NULL,
        content TEXT NOT NULL,
        detail TEXT,
        deadline DATE,
        status TEXT,
        type TEXT,
        done INTEGER DEFAULT 0,
        created_at DATE,
        updated_at DATE
    );

    CREATE TABLE IF NOT EXISTS cycle_todo_status(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        todo_id INTEGER NOT NULL,
        status TEXT,
        date TEXT,
        created_at DATE,
        updated_at DATE
    );

    CREATE TABLE IF NOT EXISTS users(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        username TEXT NOT NULL UNIQUE,
        avatar TEXT,
        hash TEXT NOT NULL,
        created_at DATE,
        updated_at DATE
    );

    `
	_, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}
}

func readConfg() {

	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	customConfigSrc, err := ioutil.ReadFile("config.custom.yaml")
	if err != nil { // Handle errors reading the config file
		panic(err)
	}

	err2 := viper.ReadInConfig() // Find and read the config file

	if err2 != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err2))
	}

	err3 := viper.MergeConfig(bytes.NewBuffer(customConfigSrc))
	if err3 != nil { // Handle errors reading the config file
		panic(err3)
	}
}

func main() {
	readConfg()

	db := database.InitDB("storage.sqlite3?parseTime=true&cache=shared&mode=rwc")
	defer db.Close()

	migrate(db)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/upload", "upload")

	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello my firend")
	})
	e.POST("/signin", handlers.SignIn(db))
	e.POST("/signup", handlers.SignUp(db))
	e.GET("/captcha/*", echo.WrapHandler(captcha.Server(captcha.StdWidth, captcha.StdHeight)))
	e.POST("/captcha", handlers.GenCaptcha(db))

	r := e.Group("/auth")
	r.Use(middleware.JWT([]byte("secret")))

	r.GET("/todo", handlers.GetTodos(db))
	r.GET("/user/:userId", handlers.GetUserInfo(db))
	r.POST("/todo", handlers.PostTodo(db))
	r.DELETE("/todo/:id", handlers.DeleteTodo(db))
	r.PUT("/todo/:id", handlers.PutTodo(db))
	r.POST("/todo/:id/cycle", handlers.MarkCycleTodo(db))

	r.POST("/avatar", handlers.PostAvatar(db))
	r.POST("/avatar/base64", handlers.PostAvatarByBase64(db))

	fmt.Println("jellyfish serve on http://0.0.0.0:8000")
	e.Logger.Fatal(e.Start("0.0.0.0:8000"))
}
