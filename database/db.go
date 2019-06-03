package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func InitDB(filepath string) *sql.DB {
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

func Migrate(db *sql.DB) {
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

type Instance struct {
	DB *sql.DB
}

var dbInstance *Instance

func connectPostgres() *sql.DB {
	connStr := "postgres://" + "pqgotest:password@localhost/pqgotest?sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func GetDB() *sql.DB {
	if dbInstance == nil {
		dbInstance = new(Instance)
		dbInstance.DB = connectPostgres()
	}
	return dbInstance.DB
}
