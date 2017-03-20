package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Index useless println
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

//CatsIndex send a json back with all the cats database
func CatsIndex(w http.ResponseWriter, r *http.Request) {
	cats := Cats{
		Cat{ID: "1"},
		Cat{ID: "2"},
		Cat{ID: "3"},
	}
	log.Println("Encoding to json....")
	json.NewEncoder(w).Encode(cats)
}

//CatShow show a specific required beautiful cat
func CatShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["catId"]
	fmt.Fprintln(w, "cat show:", todoID)
}
