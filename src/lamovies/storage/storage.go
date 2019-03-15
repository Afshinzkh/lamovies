package storage

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "asdfghjk"
	dbname   = "movies_db"
)

var db *sql.DB

// MoviesStorage handles the connection to the database
type MoviesStorage struct{}

func (m *MoviesStorage) Connect() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

func (m *MoviesStorage) GetAll() {}

func (m *MoviesStorage) GetByID() {}

func (m *MoviesStorage) Add() {}
