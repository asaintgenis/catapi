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
	cat, err := searchImage("")
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(cat); err != nil {
			panic(err)
		}
		fmt.Fprintln(w)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(cat); err != nil {
		panic(err)
	}
	fmt.Fprintln(w)
}

//CatShow show a specific required beautiful cat
func CatShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	catID := vars["catId"]
	cat, err := searchImage(catID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(cat); err != nil {
			panic(err)
		}
		fmt.Fprintln(w)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(cat); err != nil {
		panic(err)
	}
	fmt.Fprintln(w)
}
