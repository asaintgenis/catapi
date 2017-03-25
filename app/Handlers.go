package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//Index useless println
func Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Welcome!")
}

//CatsIndex send a json back with all the cats database
func CatsIndex(w http.ResponseWriter, r *http.Request) {
	cat := searchImage("")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(cat); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, cat)
}

//CatShow show a specific required beautiful cat
func CatShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	catID := vars["catId"]
	cat := searchImage(catID)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(cat); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, cat)
}
