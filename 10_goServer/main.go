package main

import (
	"./people"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"os"
)

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
	fmt.Println("Hit root endpoint")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", root)
	router.HandleFunc("/people", people.ReturnAllPeopleFromDB)
	router.HandleFunc("/people/{id}", people.ReturnSinglePerson)
	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":3001", nil))
}

func main() {
	seedDB()
	handleRequests()
}

func seedDB() {
	os.Remove("./test.db")

	database, _ := sql.Open("sqlite3", "./test.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstName TEXT, lastName TEXT)")
	statement.Exec()
	statement, _ = database.Prepare("INSERT INTO people (firstName, lastName) VALUES (?, ?)")
	statement.Exec("Ryan", "McCombe")
	statement.Exec("Hannah", "Morrison")
	statement.Exec("David", "McDatabase")
}