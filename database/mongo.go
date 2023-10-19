package database

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var db *sql.DB

func GetDB() (*sql.DB, error) {
	if db == nil {
		return nil, errors.New("error open db")
	}
	return db, nil
}

func StartDB() error {
	uri := os.Getenv("DB_URI")
	if uri == "" {
		return errors.New("you must set your 'DB_URI' environmental variable")
	}

	var err error
	db, err = sql.Open("mysql", uri)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return errors.New("dont derived connections")
	}

	return nil
}

func CloseDB() error {
	if db != nil {
		err := db.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
