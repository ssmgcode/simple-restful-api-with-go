package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
}

type Articles []Article

func allArticles(rw http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test Title", Description: "Test Description", Content: "Test Content"},
	}
	fmt.Println("Endpoint Hit: All Articles Endpoint")
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(articles)
}

func testPostArticles(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Test POST")
}

func homePage(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Homepage Endpoint Hit")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage)
	router.HandleFunc("/articles", allArticles).Methods("GET")
	router.HandleFunc("/articles", testPostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":5000", router))
}

func main() {
	handleRequests()
}
