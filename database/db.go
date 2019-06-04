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
