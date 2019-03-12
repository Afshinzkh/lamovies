package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"lamovies/output"
	"lamovies/types"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// a sample struct for noDatabase scenario
var movies []types.Movie

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "asdfghjk"
	dbname   = "movies_db"
)

func init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

}

func routers() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/add", addMovie).Methods("POST")
	r.HandleFunc("/getall", getAllMovies).Methods("GET")
	r.HandleFunc("/get/{id}", getMovie).Methods("GET")

	return r
}

func main() {
	r := routers()

	fmt.Println("Listening on port 3000 ...")
	log.Fatal(http.ListenAndServe(":3000", r))
}

func addMovie(w http.ResponseWriter, r *http.Request) {
	var movie types.Movie

	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		output.Error(w, http.StatusBadRequest, err.Error())
	}

	movies = append(movies, movie)

	output.JSON(w, http.StatusOK, movie)

}

func getAllMovies(w http.ResponseWriter, r *http.Request) {
	output.JSON(w, http.StatusOK, movies)

}

func getMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, movie := range movies {
		if movie.ID == params["id"] {
			output.JSON(w, http.StatusOK, movie)
			return
		}
	}
	output.Error(w, http.StatusNotFound, "movie with this ID is not found")

}
