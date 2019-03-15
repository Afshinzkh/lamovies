package storage

import (
	"database/sql"
	"fmt"
	"lamovies/types"
	"time"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "asdfghjk"
	dbname   = "movies_db"
	table    = "movies"
)

// MoviesStorage handles the connection to the database
type MoviesStorage struct {
	db *sql.DB
}

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
	m.db = createTable(db)
	fmt.Println("Successfully connected!")
}

func (m *MoviesStorage) GetAll() (movies []types.Movie, err error) {
	rows, err := m.db.Query("SELECT * FROM movies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movie types.Movie
	for rows.Next() {
		err = rows.Scan(&movie.ID, &movie.Name, &movie.Status, &movie.DateAdded)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}

	return movies, nil
}

func (m *MoviesStorage) GetByID(id int) (types.Movie, error) {
	statement := `SELECT * FROM movies WHERE mid=$1`

	var movie types.Movie
	if err := m.db.QueryRow(statement, id).Scan(&movie.ID, &movie.Name, &movie.Status, &movie.DateAdded); err != nil {
		return types.Movie{}, err
	}

	return movie, nil
}

func (m *MoviesStorage) Add(movie types.Movie) (types.Movie, error) {
	now := time.Now()
	statement := `INSERT INTO movies (name, status, date_added) 
				  VALUES ($1, $2, $3)
				  Returning mid`
	id := 0
	err := m.db.QueryRow(statement, movie.Name, movie.Status, now).Scan(&id)
	if err != nil {
		return types.Movie{}, err
	}
	movie.ID = id
	movie.DateAdded = now
	return movie, nil
}

func createTable(db *sql.DB) *sql.DB {
	statement := ` CREATE TABLE IF NOT EXISTS "movies" (
		mid BIGSERIAL PRIMARY KEY NOT NULL,
		name TEXT NOT NULL,
		status VARCHAR(20) NOT NULL,
		date_added TIMESTAMP)`

	_, err := db.Exec(statement)
	if err != nil {
		panic(err)
	}

	return db
}
