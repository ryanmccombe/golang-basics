package main

import (
	"./articles"
	"fmt"
	"log"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
	fmt.Println("Hit root endpoint")
}

func handleRequests() {
	http.HandleFunc("/", root)
	http.HandleFunc("/all", articles.ReturnAllArticles)
	log.Fatal(http.ListenAndServe(":3001", nil))
}

func main() {
	handleRequests()
}
