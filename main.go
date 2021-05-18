package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rest-api-in-go/static"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test Title", Desc: "Test description", Content: "Hello World"},
	}
	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func getGopher(w http.ResponseWriter, r *http.Request) {
	w.Write(static.Gopher)
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	http.HandleFunc("/gopher", getGopher)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
func main() {
	handleRequest()
}
