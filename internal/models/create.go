package models

import (
	"database/sql"
	"io/ioutil"
)

func CreateDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	filebyte, err := ioutil.ReadFile("./internal/models/statements.sql")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(string(filebyte))
	if err != nil {
		return nil, err
	}
	return db, nil
}
