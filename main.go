package main

import (
	"database/sql"
	"go-echo-vue/handlers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
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
    CREATE TABLE IF NOT EXISTS todo(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        content TEXT NOT NULL,
        detail TEXT,
        dealine DATE,
        status: TEXT,
        created_at DATE,
        updated_at DATE
    );

    CREATE TABLE IF NOT EXISTS keep(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        content TEXT NOT NULL,
        status: TEXT,
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

	e.GET("/api/tasks", handlers.GetTasks(db))
	e.PUT("/api/tasks", handlers.PutTask(db))
	e.DELETE("/tasks/:id", handlers.DeleteTask(db))

	e.Run(standard.New(":8000")) // Start as a web server
	e.Run(standard.New(":8000"))
}
