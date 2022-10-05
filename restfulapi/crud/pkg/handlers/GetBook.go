package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"crud/pkg/mocks"
	"github.com/gorilla/mux"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
  // Read dynamic id parameter
  vars := mux.Vars(r)
  id, _ := strconv.Atoi(vars["id"])

  // Iterate over all the Mock books
  for _, book := range mocks.Books {
	if book.Id == id {
	  // If ids are equal send book as response
	  w.WriteHeader(http.StatusOK)
	  w.Header().Add("Content-Type", "application/json")
	  json.NewEncoder(w).Encode(book)
	  break
	}
  }
}