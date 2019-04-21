package main

import (
	"./articles"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
	fmt.Println("Hit root endpoint")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", root)
	router.HandleFunc("/all", articles.ReturnAllArticles)
	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":3001", nil))
}

func main() {
	handleRequests()
}
