package main

import (
	"database/sql"

	"jellyfish/database"
	"jellyfish/handlers"

	"fmt"
	"github.com/dchest/captcha"

	_ "github.com/labstack/gommon/log"

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
        done INTEGER DEFAULT 0,
        created_at DATE,
        updated_at DATE
    );

    CREATE TABLE IF NOT EXISTS users(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        username TEXT NOT NULL UNIQUE,
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

func main() {
	db := database.InitDB("storage.sqlite3?parseTime=true&cache=shared&mode=rwc")
	defer db.Close()

	migrate(db)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/signin", handlers.SignIn(db))
	e.POST("/signup", handlers.SignUp(db))
	e.GET("/captcha/*", echo.WrapHandler(captcha.Server(captcha.StdWidth, captcha.StdHeight)))
	e.POST("/captcha", handlers.GenCaptcha(db))

	r := e.Group("/auth")
	r.Use(middleware.JWT([]byte("secret")))

	r.GET("/auth/todo", handlers.GetTodos(db))
	r.POST("/auth/todo", handlers.PostTodo(db))
	r.DELETE("/auth/todo/:id", handlers.DeleteTodo(db))
	r.PUT("/auth/todo/:id", handlers.PutTodo(db))

	fmt.Println("jellyfish serve on http://0.0.0.0:8000")
	e.Logger.Fatal(e.Start("0.0.0.0:8000"))
}
