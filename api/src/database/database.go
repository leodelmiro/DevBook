package database

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	db, dbError := sql.Open("mysql", config.DbConnection)
	if dbError != nil {
		return nil, dbError
	}

	if dbError = db.Ping(); dbError != nil {
		db.Close()
		return nil, dbError
	}

	return db, nil
}
