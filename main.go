package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Book struct {
	Title       string `json:"Title"`
	Description string `json:"desc"`
	Author      string `json:"author"`
}

// let's declare a global Books array
// that we can then populate in our main function
// to simulate a database
var Books []Book

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	Books = []Book{
		Book{Title: "This Side of Paradise",
			Description: "The book examines the lives and morality of carefree American youth at the dawn of the Jazz Age. Its protagonist Amory Blaine is an attractive middle-class student at Princeton University who dabbles in literature and engages in a series of romances with flappers. The novel explores the theme of love warped by greed and status seeking, and takes its title from a line of Rupert Brooke's poem Tiare Tahiti.",
			Author:      "Article Content",
		},
		Book{Title: "Dracula",
			Description: "As an epistolary novel, the narrative is related through letters, diary entries, and newspaper articles. It has no single protagonist, but opens with solicitor Jonathan Harker taking a business trip to stay at the castle of a Transylvanian noble, Count Dracula.",
			Author:      "Bram Stoker",
		},
	}
	handleRequests()
}

func returnAllBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Books)
}
