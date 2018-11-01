package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Article holds the api information
type Article struct {
	ID      int    `json:"ID"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// Articles is an splice of Articles
type Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	fmt.Println("Endpoint Hit: returnAllArticles")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Fprintf(w, "Key: "+key)
}

func testPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test POST Endpoint worked!")
}

func handleRequests() {
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/api/", homePage)
	myRouter.HandleFunc("/api/articles", returnAllArticles).Methods("GET")
	myRouter.HandleFunc("/api/articles", testPost).Methods("POST")
	myRouter.HandleFunc("/api/articles/{id}", returnSingleArticle).Methods("GET")
	log.Fatal(http.ListenAndServe(":5000", myRouter))
}

func main() {
	handleRequests()
}
