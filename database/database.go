package database

import (
	"database/sql"
	"github.com/juju/errors"
	_ "github.com/lib/pq"
	configs "jellyfish/config"
)

type AppDataSource struct {
	RDS *sql.DB
}

func connectDatabase(rdsConfig configs.RDSConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", rdsConfig.DatabaseUrl)
	if err != nil {
		return nil, errors.Trace(err)
	}
	return db, nil
}

func GetDatabase(dataSourceConfig configs.DataSourceConfig) (*AppDataSource, error) {
	db, err := connectDatabase(dataSourceConfig.RDS)
	if err != nil {
		return nil, err
	}
	return &AppDataSource{RDS: db}, nil
}
