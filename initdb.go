package main

import (
	"database/sql"
)

func initdb(DataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("postgres", DataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
