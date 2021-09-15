package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

func homePage(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Homepage Endpoint Hit")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func main() {
	handleRequests()
}
