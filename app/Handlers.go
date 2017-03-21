package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/asaintgenis/catapi/data/cat"
)

//Index useless println
func Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Welcome!")
}

//CatsIndex send a json back with all the cats database
func CatsIndex(w http.ResponseWriter, r *http.Request) {
	cats := cat.Cats{
		cat.Cat{ID: "1"},
		cat.Cat{ID: "2"},
		cat.Cat{ID: "3"},
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(cats); err != nil {
		panic(err)
	}
}

//CatShow show a specific required beautiful cat
func CatShow(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//todoID := vars["catId"]
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintln(w, "Not yet Implemented !")
}
