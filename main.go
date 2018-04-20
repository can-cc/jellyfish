package main

import (
	"database/sql"

	"jellyfish/handlers"

	"fmt"

	_ "github.com/labstack/gommon/log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"

	_ "github.com/mattn/go-sqlite3"
)

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	// Here we check for any db errors then exit
	if err != nil {
		panic(err)
	}

	// If we don't get any errors but somehow still don't get a db connection
	// we exit as well
	if db == nil {
		panic("db nil")
	}
	return db
}

func migrate(db *sql.DB) {
	sql := `
    CREATE TABLE IF NOT EXISTS todos(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
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
    `
	_, err := db.Exec(sql)
	// Exit if something goes wrong with our SQL statement above
	if err != nil {
		panic(err)
	}
}

func main() {
	// Create a new instance of Echo
	db := initDB("storage.sqlite3")
	migrate(db)

	e := echo.New()

	e.Use(middleware.Logger())

	e.GET("/todo", handlers.GetTodos(db))
	e.POST("/todo", handlers.PostTodo(db))
	e.DELETE("/todo/:id", handlers.DeleteTodo(db))

	fmt.Printf("jellyfish serve on http://localhost:8000")
	e.Run(standard.New(":8000")) // Start as a web server
}
