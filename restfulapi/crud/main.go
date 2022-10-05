package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Book struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

var Books = []Book{
	{
		Id:     1,
		Title:  "Golang",
		Author: "Gopher",
		Desc:   "A book for Go",
	},
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Books)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/books", GetAllBooks)

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}
