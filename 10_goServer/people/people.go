package people

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Person struct {
	Id int `json:"Id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}

type People []Person

func ReturnAllPeople(w http.ResponseWriter, r *http.Request){
	people := People{
		Person{Id: 1, FirstName: "Ryan", LastName: "McCombe"},
		Person{Id: 2, FirstName: "Hannah", LastName: "Morrison"},
	}
	fmt.Println("Endpoint Hit: returnAllPeople")

	json.NewEncoder(w).Encode(people)
}

func ReturnSinglePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	fmt.Fprintf(w, "Key: " + key)
}

func ReturnAllPeopleFromDB(w http.ResponseWriter, r *http.Request) {
	// TODO: keep the DB conn open
	database, _ := sql.Open("sqlite3", "./test.db")
	rows, _ := database.Query("SELECT id, firstName, lastName FROM people")
	people := People{}
	var id int
	var firstName string
	var lastName string
	for rows.Next() {
		rows.Scan(&id, &firstName, &lastName)
		people = append(people, Person{
			Id: id,
			FirstName: firstName,
			LastName: lastName,
		})

		// fmt.Println(strconv.Itoa(id) + ": " + firstName + " " + lastName)
	}
	json.NewEncoder(w).Encode(people)
}