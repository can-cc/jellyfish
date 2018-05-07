package main

import (
	"database/sql"

	"jellyfish/database"
	"jellyfish/handlers"

	"fmt"

	_ "github.com/labstack/gommon/log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
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
        created_at DATE,
        updated_at DATE
    );

    CREATE TABLE IF NOT EXISTS keeps(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        content TEXT NOT NULL,
        status TEXT,
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
	// Exit if something goes wrong with our SQL statement above
	if err != nil {
		panic(err)
	}
}

func main() {
	db := database.InitDB("storage.sqlite3?parseTime=true")
	migrate(db)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/signin", handlers.SignIn(db))

	r := e.Group("")
	r.Use(middleware.JWT([]byte("secret")))

	r.GET("/todo", handlers.GetTodos(db))
	r.POST("/todo", handlers.PostTodo(db))
	r.DELETE("/todo/:id", handlers.DeleteTodo(db))

	fmt.Println("jellyfish serve on http://localhost:8000")
	e.Run(standard.New("0.0.0.0:8000")) // Start as a web server
}
