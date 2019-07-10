package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Instance struct {
	DB *sql.DB
}

var dbInstance *Instance

func connectDB() *sql.DB {
	DBNAME := viper.GetString("DBNAME")
	DBHOST := viper.GetString("DBHOST")
	DBUSER := viper.GetString("DBUSER")
	DBPASSWORD := viper.GetString("DBPASSWORD")

	connStr := "postgres://" + DBUSER + ":" + DBPASSWORD + "@" + DBHOST + ":5432/" + DBNAME + "?sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func GetDB() *sql.DB {
	if dbInstance == nil {
		dbInstance = new(Instance)
		dbInstance.DB = connectDB()
	}
	return dbInstance.DB
}

func CheckDBConnect() {
	db := GetDB()
	err := db.Ping()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
